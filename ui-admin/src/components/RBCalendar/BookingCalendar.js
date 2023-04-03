import { useCallback, useEffect, useState } from 'react';

import moment from 'moment';
import { Calendar, momentLocalizer } from 'react-big-calendar';
import 'react-big-calendar/lib/css/react-big-calendar.css';

import '../../assets/style/calendar.css';
import useCheckMobileScreen from '../../hooks/useCheckMobileScreen';
import {
  defaultMinTime,
  defaultMaxTime,
  rbcLocale,
  startEventTitleWeek,
  calendarViews,
  startAccessor,
  endAccessor,
  messages,
  formats,
} from './CalendarHelpers';

moment.locale(rbcLocale, startEventTitleWeek);

function BookingCalendar({ date, events, onNavigate, onSelectEvent }) {
  const initialCalendarView = useCheckMobileScreen() ? 'week' : 'month';
  const [calendarView, setCalendarView] = useState(initialCalendarView);
  const [minTime, setMinTime] = useState(defaultMinTime);
  const [maxTime, setMaxTime] = useState(defaultMaxTime);

  const handleCalendarView = (newView) => {
    setCalendarView(newView);
  };

  const onEventPropGetter = (e) => ({
    ...(e && e?.bkng && { className: 'rbc-event-booked' }),
  });

  const titleAccessor = ({ start, bkng = {} }) => {
    const startTm = new Date(start);
    const hour = startTm.getHours();

    // If min is smaller than 10 then add an extra 0 before the min digit
    let min = startTm.getMinutes();
    if (min < 10) { min = `0${min}`; }

    if (Object.keys(bkng).length !== 0) {
      return <span>{`${bkng.name} ${hour}:${min}`}</span>;
    }

    return <span>{`${hour}:${min}`}</span>;
  };

  const handleMinMaxTime = useCallback(() => {
    if (!events || events.length === 0) return;

    let minTm = new Date(events[0].start);
    let maxTm = new Date(events[0].end);

    events.forEach((e) => {
      const sd = new Date(e.start);
      const ed = new Date(e.end);

      if (sd.getHours() < minTm.getHours() || (sd.getHours() === minTm.getHours() && sd.getMinutes() < minTm.getMinutes())) {
        minTm = sd;
      }

      if (sd.getDay() !== ed.getDay()) {
        maxTm.setHours(23, 59);
      } else if ((maxTm.getHours() !== 23 && maxTm.getMinutes() !== 59) || ed.getHours() > maxTm.getHours() || (ed.getHours() === maxTm.getHours() && ed.getMinutes() > maxTm.getMinutes())) {
        maxTm = ed;
      }
    });

    setMinTime(minTm);
    setMaxTime(maxTm);
  }, [events]);

  useEffect(() => {
    handleMinMaxTime();
  }, [events, handleMinMaxTime]);

  return (
    <div className="rbc-container">

      <Calendar
        key="rbc"
        // ------ Localize settings ------
        date={date}
        localizer={momentLocalizer(moment)}
        // ------ View settings ------
        formats={formats}
        view={calendarView}
        views={calendarViews}
        onView={handleCalendarView}
        // ------ Event settings ------
        events={events}
        endAccessor={endAccessor}
        startAccessor={startAccessor}
        titleAccessor={titleAccessor}
        onNavigate={onNavigate}
        onSelectEvent={onSelectEvent}
        eventPropGetter={onEventPropGetter}
        // ------ Slot settings ------
        popup
        step={15}
        timeslots={4}
        min={minTime}
        max={maxTime}
        messages={messages}
      />

    </div>
  );
}

export default BookingCalendar;
