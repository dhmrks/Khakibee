const rbcLocale = 'en-GB';
const startEventTitleWeek = { week: { dow: 1, doy: 1 } }; // Start event title week from Monday

const weekStartDate = '2022/08/01';
const defaultMinTime = '10:00';
const defaultMaxTime = '23:59';
const calendarViews = ['month', 'week', 'day'];

const startAccessor = ({ start }) => new Date(start);
const endAccessor = ({ start, end }) => {
  const sd = new Date(start);
  const ed = new Date(end);

  // if event ends the next day, end at 00:00:00 of the same day
  if (sd.getDay() !== ed.getDay()) {
    sd.setHours(23, 59, 59);
    return sd;
  }

  return ed;
};
const messages = {
  today: <i className="bi bi-calendar-check" />,
  next: <i className="bi bi-caret-right-fill" />,
  previous: <i className="bi bi-caret-left-fill" />,
  noEventsInRange: 'There are no events in this range.',
};
const formats = {
  dayFormat: 'DD dd',
  timeGutterFormat: (date, culture, localizer) => localizer.format(date, 'H:mm', culture),
  eventTimeRangeFormat: () => '',
};

export {
  defaultMinTime,
  defaultMaxTime,
  rbcLocale,
  startEventTitleWeek,
  calendarViews,
  startAccessor,
  endAccessor,
  messages,
  formats,
  weekStartDate,
};
