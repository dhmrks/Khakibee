import { useCallback, useEffect, useState } from 'react';

import service from '../../hooks/service';
import { formatDate, addMinutes, getWeekDays, generateTimestampID, isEmptyObj, isValidDate } from '../../utils/misc';
import ErrorHandler from '../ErrorHandler';
import { weekStartDate } from '../RBCalendar/CalendarHelpers';
import ScheduleCalendar from '../RBCalendar/ScheduleCalendar';
import Toast from '../UI/Toast';
import statuses from './ScheduleStatuses';

function Schedule({ gid, dur }) {
  const [events, setEvents] = useState();
  const [activeBy, setActiveBy] = useState();
  const [scheduleType, setScheduleType] = useState('current');

  const { data: schedule, error, request: requestSchedule } = service.schedule.useAll();
  const { ok: createOk, request: createSchedule } = service.schedule.useCreate();
  const { ok: editOk, request: editSchedule } = service.schedule.useUpdate();

  const handleScheduleTypeAndDate = (name) => {
    const value = name === 'upcoming' ? schedule?.upcoming?.active_by : null;

    setActiveBy(value);
    setScheduleType(name);
  };

  const save = (schdl, dt) => {
    const editedSchedule = {};

    schdl.forEach((event) => {
      if (event.status === 0) return;

      const startTime = event.start;
      const weekDay = startTime.getDay();

      let minutes = startTime.getMinutes();
      if (minutes < 10) minutes = `0${minutes}`;
      const eventTime = `${startTime.getHours()}:${minutes}`;

      const hasKey = weekDay in editedSchedule;
      if (!hasKey) editedSchedule[weekDay] = [];

      editedSchedule[weekDay].push(eventTime);
    });

    const args = {
      params: { gid },
      body: editedSchedule,
    };

    if (scheduleType === 'upcoming') {
      if (isValidDate(dt)) args.body.active_by = formatDate(dt);
      createSchedule(args);
      return;
    }

    editSchedule(args);
  };

  const handleScheduleEvents = useCallback((s) => {
    const evnts = [];
    if (isEmptyObj(s)) return evnts;

    const weekDays = getWeekDays(weekStartDate);

    for (let i = 0; i <= 6; i += 1) {
      if (!s[i]) {
        // eslint-disable-next-line no-continue
        continue;
      }

      const dt = weekDays[i];

      for (let j = 0; j < s[i].length; j += 1) {
        const [hours, minutes] = s[i][j].split(':');
        dt.setHours(hours);
        dt.setMinutes(minutes);

        const newEvent = {
          id: generateTimestampID(dt),
          start: new Date(dt),
          end: addMinutes(dur, dt),
          status: statuses.saved,
        };
        evnts.push(newEvent);
      }
    }

    return evnts;
  }, [dur]);

  useEffect(() => {
    const args = { params: { gid } };
    requestSchedule(args);
  }, [gid, editOk, createOk, requestSchedule]);

  useEffect(() => {
    let schdlEvents = [];

    if (scheduleType === 'current') {
      schdlEvents = handleScheduleEvents(schedule);
    } else if (scheduleType === 'upcoming') {
      const upcomingSchdl = schedule?.upcoming && Object.keys(schedule?.upcoming).length > 1 ? schedule.upcoming : {};
      schdlEvents = handleScheduleEvents(upcomingSchdl);
    }

    setEvents(schdlEvents);
  }, [schedule, scheduleType, handleScheduleEvents]);

  const errorModal = error && <ErrorHandler error={error} />;
  const successToast = (editOk || createOk) && <Toast />;

  return (
    <>
      {errorModal}
      {successToast}

      <div className="row mx-1">
        <div className="d-flex justify-content-center">

          <div className="btn-group" role="group">
            <input type="radio" className="btn-check" name="current" id="current" autoComplete="off" checked={scheduleType === 'current'} onChange={(e) => handleScheduleTypeAndDate(e.target.name)} />
            <label className="btn btn-sm btn-outline-primary" htmlFor="current">Current</label>

            {!isEmptyObj(schedule) && (
              <>
                <input type="radio" className="btn-check" name="upcoming" id="upcoming" autoComplete="off" checked={scheduleType === 'upcoming'} onChange={(e) => handleScheduleTypeAndDate(e.target.name)} />
                <label className="btn btn-sm btn-outline-primary" htmlFor="upcoming">{isEmptyObj(schedule?.upcoming) ? 'Add new' : 'Upcoming'}</label>
              </>
            )}

          </div>

        </div>
      </div>

      <ScheduleCalendar
        key={scheduleType}
        dur={dur}
        date={weekStartDate}
        schedule={events}
        onSave={save}
        scheduleType={scheduleType}
        activeBy={activeBy}
      />

    </>
  );
}

export default Schedule;
