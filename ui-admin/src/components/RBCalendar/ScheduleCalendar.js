import { useCallback, useState, forwardRef, useEffect } from 'react';

import classnames from 'classnames';
import moment from 'moment';
import { Calendar, momentLocalizer } from 'react-big-calendar';
import 'react-big-calendar/lib/css/react-big-calendar.css';
import withDragAndDrop from 'react-big-calendar/lib/addons/dragAndDrop';
import DatePicker from 'react-datepicker';
import TimePicker from 'react-time-picker';
import 'react-datepicker/dist/react-datepicker.css';
import 'react-big-calendar/lib/addons/dragAndDrop/styles.css';

import '../../assets/style/calendar.css';
import { addWeeks, addMinutes, isEventPermited, deleteObjFromArray, generateTimestampID } from '../../utils/misc';
import statuses from '../Schedule/ScheduleStatuses';
import Modal from '../UI/Modal';
import {
  defaultMinTime,
  defaultMaxTime,
  rbcLocale,
  startEventTitleWeek,
  startAccessor,
  endAccessor, formats,
  weekStartDate,
} from './CalendarHelpers';

const calendarView = 'week';
moment.locale(rbcLocale, startEventTitleWeek);

const DnDCalendar = withDragAndDrop(Calendar);

function scheduleToolbar() {
  return (<> </>);
}

function ScheduleCalendar({ dur, date, schedule, onSave, scheduleType, activeDate }) {
  const [events, setEvents] = useState(schedule);
  const [selectedEvent, setSelectedEvent] = useState();
  const [showRemoveEvent, setShowRevomeEvent] = useState(false);

  const [minTime, setMinTime] = useState(defaultMinTime);
  const [maxTime, setMaxTime] = useState(defaultMaxTime);

  const isUpcomingSchedule = scheduleType === 'upcoming';
  const [activeBy, setActiveBy] = useState(activeDate ? new Date(activeDate) : addWeeks(3));

  const timeToolsClassnames = classnames('d-flex', 'justify-content-start');
  const saveBtnClassnames = classnames('d-flex', 'justify-content-end', { 'mt-2 mt-sm-0': isUpcomingSchedule });

  useEffect(() => {
    setEvents(schedule);
  }, [schedule]);

  const onSelectEvent = (event) => setSelectedEvent(event);
  const handleShowRemoveEvent = (value) => setShowRevomeEvent(value);

  const onEventPropGetter = (e) => {
    let statusClass = '';
    switch (e.status) {
      case (statuses.new):
        statusClass = 'rbc-schdl-new';
        break;
      case (statuses.deleted):
        statusClass = 'rbc-schdl-deleted';
        break;
      default:
        statusClass = 'rbc-schdl-saved';
    }

    return ({ ...({ className: statusClass }) });
  };

  const handleResetEvent = (id) => {
    let updatedEvents = [];
    setEvents((prev) => {
      updatedEvents = prev.map((e) => {
        if (e.id === id) return { ...e, status: statuses.saved };
        return e;
      });

      return updatedEvents;
    });
    return updatedEvents;
  };

  const onSelectSlot = useCallback((event) => {
    const startTime = event.start;
    const endTime = addMinutes(dur, startTime);

    if (!isEventPermited(event.id, startTime, endTime, events)) return;

    const newEvent = {
      id: generateTimestampID(startTime),
      start: startTime,
      end: endTime,
      status: statuses.new,
    };

    setEvents((prev) => [...prev, newEvent]);
  }, [dur, events, setEvents]);

  const onEventDrop = useCallback(({ event, start }) => {
    const end = addMinutes(dur, start);
    const id = generateTimestampID(start);

    if (!isEventPermited(event.id, start, end, events)) return;

    const index = events.findIndex((e) => e.id === id);
    if (id === generateTimestampID(event.id) && index >= 0) return;
    if (index >= 0) {
      let updatedEvents = handleResetEvent(id);

      updatedEvents = deleteObjFromArray(event.id, updatedEvents);
      setEvents(updatedEvents);
      return;
    }

    setEvents((prev) => {
      let updatedEvents = [];

      if (event.status === statuses.saved) {
        updatedEvents = prev.map((e) => {
          if (e.id === event.id) return { ...e, status: statuses.deleted };
          return e;
        });
        updatedEvents = [...updatedEvents, { id, start, end, status: statuses.new }];
      } else {
        updatedEvents = deleteObjFromArray(event.id, events);
        updatedEvents = [...updatedEvents, { id, start, end, status: statuses.new }];
      }

      return updatedEvents;
    });
  }, [dur, events, setEvents]);

  const onEventDelete = useCallback(() => {
    setEvents((prev) => {
      let updatedEvents = [];

      if (selectedEvent.status === statuses.new) {
        updatedEvents = prev.filter((e) => e.id !== selectedEvent.id);
      } else {
        updatedEvents = prev.map((e) => {
          if (e.id === selectedEvent.id) return { ...e, status: statuses.deleted };
          return e;
        });
      }

      return updatedEvents;
    });

    setSelectedEvent();
    setShowRevomeEvent(false);
  }, [selectedEvent]);

  const titleAccessor = ({ id, start, status }) => {
    const startTm = new Date(start);
    const end = addMinutes(dur, start);
    const hour = startTm.getHours();

    let min = startTm.getMinutes(); // If sec is smaller than 10 then add an extra 0 before the sec digit
    if (min < 10) min = `0${min}`;

    return (
      <div className="d-flex flex-sm-row flex-column justify-content-md-between justify-content-sm-center ms-1 ms-sm-0">
        <span className="mt-2 ms-1 ms-sm-0">{`${hour}:${min}`}</span>

        {!status && isEventPermited(id, start, end, events) && <button type="button" className="btn btn-sm me-2 me-sm-0" onClick={() => handleResetEvent(id)}> <i className="bi bi-arrow-90deg-left text-white" /></button>}
        {status > 0 && <button type="button" className="btn btn-sm me-3 me-sm-0" onClick={() => handleShowRemoveEvent(true)}> <i className="bi bi-trash text-white" /></button>}

      </div>
    );
  };

  const modalData = {
    header: 'Remove schedule event',
    body: <>Are you sure you want to remove the existing event from the schedule ?</>,
    icon: 'bi bi-trash-fill',
    submitAction: 'Remove',
    cancelAction: 'Back',
  };
  const removeEventModal = showRemoveEvent && <Modal data={modalData} onSubmitAction={onEventDelete} onCancelAction={handleShowRemoveEvent} />;

  const handleMinTimePicker = (tm) => {
    let time = tm;
    if (tm === null) time = defaultMinTime;

    setMinTime(time);
  };

  const handleMaxTimePicker = (tm) => {
    let time = tm;
    if (tm === null) time = defaultMaxTime;

    setMaxTime(time);
  };

  const onResetSchedule = () => {
    setEvents(schedule);
    setMinTime(defaultMinTime);
    setMaxTime(defaultMaxTime);
  };

  const draggableAccessor = (e) => {
    let isDraggable = false;
    if (e.status > 0) isDraggable = true;

    return isDraggable;
  };

  return (
    <>
      {removeEventModal}

      <div className="mx-1">

        <div className="row mt-2">

          <div className="col-9">
            <div className={timeToolsClassnames}>

              <div className="mx-1">
                <label className="me-2" htmlFor="minTime">Min: </label>
                <TimePicker
                  value={minTime}
                  clockIcon={null}
                  clearIcon={null}
                  className="custom-time-picker"
                  minTime={defaultMinTime}
                  maxTime={maxTime}
                  hourPlaceholder="HH"
                  minutePlaceholder="mm"
                  format="HH:mm"
                  onChange={(tm) => handleMinTimePicker(tm)}
                />
              </div>

              <div className="mx-1">
                <label className="me-2" htmlFor="maxTime">Max: </label>
                <TimePicker
                  value={maxTime}
                  clockIcon={null}
                  clearIcon={null}
                  className="custom-time-picker"
                  minTime={minTime}
                  maxTime={defaultMaxTime}
                  hourPlaceholder="HH"
                  minutePlaceholder="mm"
                  format="HH:mm"
                  onChange={(tm) => handleMaxTimePicker(tm)}
                />
              </div>

              {isUpcomingSchedule && (
                <div className="mx-1">
                  <label className="me-2" htmlFor="activeBy">Active by: </label>
                  <DatePicker
                    closeOnScroll
                    dateFormat="yyyy/MM/dd"
                    minDate={addWeeks(3)}
                    maxDate={addWeeks(24)}
                    selected={activeBy}
                    onChange={(d) => setActiveBy(d)}
                    showDisabledMonthNavigation
                    customInput={<DatePickerInput />}
                  />
                </div>
              )}

            </div>
          </div>

          <div className="col-3">
            <div className={saveBtnClassnames}>
              <button type="button" className="btn btn-sm btn-link" onClick={onResetSchedule}>Reset</button>
              <button type="button" className="btn btn-sm btn-primary ms-1" onClick={() => onSave(events, activeBy)}>Save</button>
            </div>
          </div>

        </div>

        <div className="row mt-1">

          <div className="d-flex justify-content-start">

            <label htmlFor="keys" className="ms-1 me-3">Keys: </label>

            <div className="d-flex me-2">
              <div id="saved" className="square text-bg-blue mt-1" />
              <label htmlFor="saved" className="ms-1">Saved</label>
            </div>

            <div className="d-flex me-2">
              <div id="new" className="square text-bg-new mt-1" />
              <label htmlFor="new" className="ms-1">New</label>
            </div>

            <div className="d-flex">
              <div id="deleted" className="square text-bg-deleted mt-1" />
              <label htmlFor="deleted" className="ms-1">Deleted</label>
            </div>

          </div>

        </div>

        <DnDCalendar
          key="dnd-rbc"
          // ------ Localize settings ------
          defaultDate={date}
          localizer={momentLocalizer(moment)}
          // -------- View settings --------
          formats={{ ...formats, dayFormat: 'ddd' }}
          views={[calendarView]}
          defaultView={calendarView}
          // -------- Event settings -------
          events={events}
          endAccessor={endAccessor}
          startAccessor={startAccessor}
          titleAccessor={titleAccessor}
          eventPropGetter={onEventPropGetter}
          longPressThreshold={10}
          // ------ Drag&Drop settings -----
          selectable
          onEventDrop={onEventDrop}
          onSelectSlot={onSelectSlot}
          onSelectEvent={onSelectEvent}
          draggableAccessor={draggableAccessor}
          // -------- Resize settings -------
          resizable={false}
          // --------- Slot settings --------
          popup
          step={15}
          timeslots={4}
          min={new Date(`${weekStartDate} ${minTime}`)}
          max={new Date(`${weekStartDate} ${maxTime}`)}
          components={{ toolbar: scheduleToolbar }}
        />

      </div>

    </>
  );
}

const DatePickerInput = forwardRef(({ value, onClick }, ref) => (
  <input type="text" className="date-picker-input" value={value} onClick={onClick} ref={ref} readOnly />
));

export default ScheduleCalendar;
