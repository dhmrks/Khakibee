import { useCallback, useEffect, useState } from 'react';

import { Outlet, useNavigate } from 'react-router-dom';

import service from '../../hooks/service';
import { monthDiff } from '../../utils/misc';
import ErrorHandler from '../ErrorHandler';
import BookingCalendar from '../RBCalendar/BookingCalendar';

const currentDate = new Date();

function Calendar({ gid }) {
  const navigate = useNavigate();

  const [offset, setOffset] = useState(0);
  const [newDate, setNewDate] = useState();
  const [rbcDate, setRbcDate] = useState(currentDate);
  const [chooseNewDate, setChooseNewDate] = useState(false);

  const { data: events, ok, error, request: requestBookings } = service.games.useCalendarSSE();

  useEffect(() => {
    const args = {
      params: { gid },
      query: { offset },
    };
    requestBookings(args);
  }, [gid, offset, requestBookings]);

  const errorModal = error && <ErrorHandler error={error} />;

  const handleNavigation = useCallback((navDte) => {
    if (typeof ok === 'undefined') return;

    setRbcDate(navDte);
    setOffset(monthDiff(currentDate, navDte));
  }, [ok, setOffset]);

  const handleSelectEvent = useCallback((e) => {
    const isBooked = !!e?.bkng?.bkng_id;

    if (chooseNewDate) {
      if (isBooked) return;
      setNewDate(e.start);
      setChooseNewDate(false);
    } else if (isBooked) {
      navigate(`/games/${gid}/booking/${e?.bkng?.bkng_id}`);
    } else {
      setNewDate(e.start);
      navigate(`/games/${gid}/booking/create`);
    }
  }, [gid, chooseNewDate, navigate, setChooseNewDate, setNewDate]);

  const onChooseNewDate = (value) => setChooseNewDate(value);
  const resetOutletContext = () => {
    setNewDate();
    setChooseNewDate(false);
  };

  return (
    <>
      <div className="d-flex justify-content-between">
        {chooseNewDate && (
          <button type="button" className="offcanvas-arrow" onClick={() => setChooseNewDate(false)}>
            <i className="bi bi-caret-left-fill offcanvas-caret" />
          </button>
        )}
      </div>

      {errorModal}

      <BookingCalendar
        date={rbcDate}
        events={events}
        onNavigate={handleNavigation}
        onSelectEvent={handleSelectEvent}
      />

      <Outlet context={[newDate, chooseNewDate, onChooseNewDate, resetOutletContext]} />
    </>
  );
}

export default Calendar;
