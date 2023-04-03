const generateTimestampID = (d) => new Date(d).getTime();

const isEmptyObj = (obj) => {
  if (obj === undefined || obj === null) return true;

  return obj && Object.keys(obj).length === 0 && Object.getPrototypeOf(obj) === Object.prototype;
};

const addMinutes = (minutes, date = new Date()) => {
  const newDate = new Date();

  return new Date(newDate.setTime(date.getTime() + minutes * 60000));
};

const addWeeks = (numOfWeeks, date = new Date()) => date.setDate(date.getDate() + numOfWeeks * 7) && date;

const monthDiff = (d1, d2) => {
  const years = d2.getFullYear() - d1.getFullYear();
  const months = (years * 12) + (d2.getMonth() - d1.getMonth());

  return months;
};

const getWeekDays = (d) => {
  const weekDays = {};
  const curr = new Date(d);

  for (let i = 0; i <= 6; i += 1) {
    const date = new Date(curr.setDate(i + 1));

    if (i === 6) weekDays[0] = date;

    weekDays[i + 1] = date;
  }

  return weekDays;
};

const formatDate = (date) => {
  const d = new Date(date);
  let month = `${d.getMonth() + 1}`;
  let day = `${d.getDate()}`;
  const year = d.getFullYear();

  if (month.length < 2) { month = `0${month}`; }
  if (day.length < 2) { day = `0${day}`; }

  return [year, month, day].join('-');
};

const isEventPermited = (id, startTime, end, events = []) => {
  let isPermitted = true;
  if (events.length === 0) return isPermitted;

  const dateEvents = events.filter((e) => e.id !== id && e.start.getDate() === startTime.getDate());

  for (let i = 0; i < dateEvents.length; i += 1) {
    isPermitted = dateEvents[i].status === 0 || dateEvents[i].start.getTime() > end.getTime() || dateEvents[i].end.getTime() < startTime.getTime();

    if (!isPermitted) break;
  }

  return isPermitted;
};

const deleteObjFromArray = (id, array) => array.filter((e) => e.id !== id);

// eslint-disable-next-line no-restricted-globals
const isValidDate = (d) => d instanceof Date && !isNaN(d);

export {
  addWeeks,
  addMinutes,
  isEmptyObj,
  monthDiff,
  formatDate,
  getWeekDays,
  isValidDate,
  isEventPermited,
  deleteObjFromArray,
  generateTimestampID,
};
