import { useEffect, useCallback } from 'react';

import { Offcanvas } from 'react-bootstrap';
import { useNavigate, useOutletContext, useParams } from 'react-router-dom';

import service from '../../hooks/service';
import ErrorHandler from '../ErrorHandler';
import BookingForm from './BookingForm';

function BookingCreate() {
  const params = useParams();
  const { gid } = params;

  const navigate = useNavigate();
  const navigateBack = useCallback(() => navigate(`/games/${gid}`), [gid, navigate]);

  const [newDate, chooseNewDate, onChooseNewDate, resetOutletContext] = useOutletContext();

  const { error, ok, request: newBooking } = service.bookings.useCreate();

  const save = (b) => {
    const args = {
      params: { gid },
      body: b,
    };
    newBooking(args);
  };

  useEffect(() => {
    if (ok) {
      navigateBack();
      resetOutletContext();
    }
  }, [ok, navigateBack, resetOutletContext]);

  const bookingContent = <BookingForm dte={newDate} onSave={save} onCancel={navigateBack} onEditDate={onChooseNewDate} />;
  const errorModal = error && <ErrorHandler error={error} />;

  return (
    <Offcanvas show={!chooseNewDate} scroll={!chooseNewDate} backdrop={!chooseNewDate} placement="end" onHide={navigateBack}>
      <Offcanvas.Header closeButton>
        <Offcanvas.Title>Create Booking</Offcanvas.Title>
      </Offcanvas.Header>
      <Offcanvas.Body>

        {errorModal}
        {bookingContent}

      </Offcanvas.Body>
    </Offcanvas>
  );
}

export default BookingCreate;
