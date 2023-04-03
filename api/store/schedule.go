package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
	"gitlab.com/khakibee/khakibee/api/model"
)

type ScheduleStore interface {
	CheckSchedule(ctx context.Context, date string, gameID int) (exists bool, err error)
	SelectScheduleByGameID(ctx context.Context, gmeID int, upcoming bool) (schdl *model.Schedule, err error)
	UpsertSchedule(ctx context.Context, gmeID int, schdl model.WeekSchedule, schdlID *int) (err error)
	UpsertNewSchedule(ctx context.Context, gmeID int, schdl model.WeekSchedule, activeBy string, schdlID *int) (err error)
}

type PSQLScheduleStore struct {
	db *sql.DB
}

func NewPSQLScheduleStore(db *sql.DB) (gameStore *PSQLScheduleStore) {
	gameStore = &PSQLScheduleStore{db}

	return
}

func (s *PSQLScheduleStore) UpsertNewSchedule(ctx context.Context, gmeID int, schdl model.WeekSchedule, activeBy string, schdlID *int) (err error) {
	queryIns := `insert into schdl (gme_id, schedule,active_by) values ($1,$2,$3)`
	queryUPd := `update schdl set schedule = $1, active_by=$2 where schdl_id=$3`
	var args []interface{}
	var query string

	schdlBT, err := json.Marshal(schdl)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"struct": schdl}).Error(err)
		return
	}

	if schdlID == nil {
		query = queryIns
		args = []interface{}{gmeID, schdlBT, activeBy}
	} else {
		query = queryUPd
		args = []interface{}{schdlBT, activeBy, schdlID}
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
	}

	return nil
}

func (s *PSQLScheduleStore) UpsertSchedule(ctx context.Context, gmeID int, schdl model.WeekSchedule, schdlID *int) (err error) {
	queryIns := `insert into schdl (gme_id, schedule,active_by) values ($1,$2,current_date)`
	queryUPd := `update schdl set schedule = $1 where schdl_id=$2`
	var args []interface{}
	var query string

	schdlBT, err := json.Marshal(schdl)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"struct": schdl}).Error(err)
		return
	}

	if schdlID == nil {
		query = queryIns
		args = []interface{}{gmeID, schdlBT}
	} else {
		query = queryUPd
		args = []interface{}{schdlBT, schdlID}
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
	}

	return nil
}

func (s *PSQLScheduleStore) CheckSchedule(ctx context.Context, date string, gameID int) (exists bool, err error) {
	query := `select schedule -> extract(dow from $2::date)::TEXT ? $3 as exists 
	from schdl 
	where gme_id = $1 and active_by <= $2::date order by active_by DESC limit 1;`

	//split date and time
	dt := strings.Split(date, " ")
	if len(dt) != 2 {
		return
	}

	args := []interface{}{gameID, dt[0], dt[1]}

	row := s.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&exists)
	if err == sql.ErrNoRows {
		err = nil
	} else if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
	}

	return
}

func (s *PSQLScheduleStore) SelectScheduleByGameID(ctx context.Context, gmeID int, upcoming bool) (schdl *model.Schedule, err error) {
	query := `Select schdl_id, schedule, dur, TO_CHAR(active_by, 'YYYY-MM-DD')
	from schdl s 
		JOIN gme g ON g.gme_id = s.gme_id 
	where s.gme_id = $1`

	condition := "\nand active_by <= CURRENT_DATE order by active_by DESC limit 1;"
	if upcoming {
		condition = "\nand active_by > CURRENT_DATE;"
	}

	query = query + condition

	args := []interface{}{gmeID}

	row := s.db.QueryRowContext(ctx, query, args...)
	var (
		wkSchdl  []byte
		dur      int
		activeBy string
		schdlID  int
	)

	err = row.Scan(&schdlID, &wkSchdl, &dur, &activeBy)
	if err == sql.ErrNoRows {
		err = nil
	} else if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query}).Error(err)
	} else {
		schdl = &model.Schedule{}
		schdl.SchdlID = schdlID
		schdl.Dur = dur
		schdl.ActiveBy = activeBy
		if err = json.Unmarshal(wkSchdl, &schdl.WkSchdl); err != nil {
			log.Logger().WithFields(logrus.Fields{"Unmarshal": "JSON error"}).Error(err)
		}
	}

	return
}
