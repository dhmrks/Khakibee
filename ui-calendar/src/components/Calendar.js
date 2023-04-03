import { useCallback, useState, useEffect } from 'react';

import withTranslation from '../HOCs/withTranslation';
import service from '../service';
import CalendarDay from './CalendarDay';

import '../assets/style/calendar.css';

const currentDate = new Date();
const currYear = currentDate.getFullYear();
const currMonth = currentDate.getMonth();
const currDay = currentDate.getDate();

function Calendar({ gameID, tr: { week, noRoomsAvail }, lng, onClick }) {
  const [showingMonth, setShowingMonth] = useState(currentDate);
  const [monthOffset, setMonthOffset] = useState(0);
  const [selectedDate, setSelectedDate] = useState(currentDate);
  const [events, setEvents] = useState();

  const { data: calendar, request: requestCalendar } = service.games.useCalendarSSE();

  useEffect(() => {
    requestCalendar({ params: { gid: gameID }, query: { offset: monthOffset } });
  }, [requestCalendar, gameID, monthOffset]);

  const displayedMonth = showingMonth.toLocaleDateString(lng, { month: 'long', year: 'numeric' });
  const displayedDate = selectedDate && selectedDate.toLocaleDateString(lng, { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' });

  const handleMonthOffset = (direction) => {
    const mo = direction === 'right' ? monthOffset + 1 : monthOffset - 1;
    // const date = new Date(currentDate.getFullYear(), currentDate.getMonth() + mo);
    setMonthOffset(mo);

    const curYear = currentDate.getFullYear();
    const curMonth = currentDate.getMonth();

    setShowingMonth(new Date(curYear, curMonth + mo, 1));
  };

  const handleLeft = () => handleMonthOffset('left');
  const handleRight = () => handleMonthOffset('right');

  const createCalendarEvents = useCallback(() => {
    if (!calendar) return {};
    const calendarEvents = {};
    calendar.forEach((event) => {
      const edt = new Date(event.start);
      const day = edt.getDate();

      if (!calendarEvents[day]) { calendarEvents[day] = []; }
      calendarEvents[day].push(event.start);
    });

    return calendarEvents;
  }, [calendar]);

  const calendarEvents = createCalendarEvents();

  const createMonthDays = useCallback(() => {
    const md = [];

    const shoYear = showingMonth.getFullYear();
    const shoMonth = showingMonth.getMonth();

    const selYear = selectedDate && selectedDate.getFullYear();
    const selMonth = selectedDate && selectedDate.getMonth();
    const selDay = selectedDate && selectedDate.getDate();

    const shoDateFirstWeekDay = new Date(shoYear, shoMonth, 1).getDay();
    const shoDateLastDay = new Date(shoYear, shoMonth + 1, 0).getDate();
    const daysInPrevMonth = new Date(shoYear, shoMonth, 0).getDate();

    for (let slot = 1; slot <= 42; slot += 1) {
      let day = slot - shoDateFirstWeekDay + 1;
      if (shoDateFirstWeekDay === 0) { day -= 7; } // if first day is sunday (0) subtract 7

      const ce = calendarEvents[day];

      if (day > shoDateLastDay && slot % 7 === 1) break;

      if (day > shoDateLastDay) { day -= shoDateLastDay; }
      if (day < 1) { day += daysInPrevMonth; }

      const today = currYear === shoYear && currMonth === shoMonth && currDay === day;
      const selected = selYear === shoYear && selMonth === shoMonth && selDay === day;

      if (!events && ce && selected && today) { setEvents(ce); } // load today events on first render

      md.push(
        <CalendarDay
          key={slot}
          day={day}
          events={ce}
          today={today}
          selected={selected}
          onClick={(dayEvents) => {
            setSelectedDate(new Date(dayEvents[0]));
            setEvents(dayEvents);
          }}
        />,
      );
    }

    return md;
  }, [showingMonth, selectedDate, events, calendarEvents]);

  const monthDays = createMonthDays();

  return (
    <>
      <div className="col-md-6">
        <div className="card">
          <div className="card-header">
            <div className="row d-flex justify-content-between">
              <div className="col-2 text-right">
                <button type="button" className="btn btn-link" disabled={monthOffset === 0} onClick={handleLeft}>
                  <i className="bi bi-arrow-left clickable" />
                </button>
              </div>
              <div className="text-center col-8"><strong>{displayedMonth}</strong></div>
              <div className="col-2">
                <button type="button" className="btn btn-link" disabled={monthOffset === 1} onClick={handleRight}>
                  <i className="bi bi-arrow-right clickable" />
                </button>
              </div>
            </div>
            <hr />
            <div className="d-flex">
              <div className="weekday-col">{week.mo}</div>
              <div className="weekday-col">{week.tu}</div>
              <div className="weekday-col">{week.we}</div>
              <div className="weekday-col">{week.th}</div>
              <div className="weekday-col">{week.fr}</div>
              <div className="weekday-col">{week.sa}</div>
              <div className="weekday-col">{week.su}</div>
            </div>
          </div>
          <div className="card-body">
            <div className="row">
              { monthDays}
            </div>
          </div>
        </div>
      </div>

      <div className="col-md-6">
        <div className="card ">
          <div className="text-center card-header">
            <strong>{displayedDate}</strong>
          </div>
          <div className="card-body d-grid gap-2">
            {!events && <div className="alert alert-danger " role="alert">{noRoomsAvail}</div>}
            {events && events.map((e) => <button type="button" key={e} className="btn btn-success btn-lg" onClick={() => onClick(e)}>{e.split(' ')[1]}</button>)}
          </div>
        </div>
      </div>
    </>
  );
}

export default withTranslation(Calendar);
