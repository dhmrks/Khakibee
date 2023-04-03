import { useCallback, useEffect, useMemo } from 'react';

import { Offcanvas } from 'react-bootstrap';
import { useNavigate, useParams, useOutletContext } from 'react-router-dom';

import service from '../../hooks/service';
import ErrorHandler from '../ErrorHandler';
import BookingForm from './BookingForm';

function Booking() {
  const params = useParams();
  const { bid, gid } = params;

  const navigate = useNavigate();
  const [newDate, chooseNewDate, onChooseNewDate, resetOutletContext] = useOutletContext();

  const { data: booking, error, request: requestBooking } = service.bookings.useByID();
  const { ok: updateOK, request: editBooking } = service.bookings.useUpdate();
  const { ok: deleteOK, request: deleteBooking } = service.bookings.useRemove();

  const navigateBack = useCallback(() => navigate(`/games/${gid}`), [gid, navigate]);

  const args = useMemo(() => ({ params: { gid, bid } }), [bid, gid]);

  const del = () => {
    deleteBooking(args);
    resetOutletContext();
  };

  const save = (b) => {
    args.body = b;
    editBooking(args);
    resetOutletContext();
  };

  useEffect(() => {
    if (updateOK || deleteOK) {
      navigateBack();
      resetOutletContext();
      return;
    }

    requestBooking(args);
  }, [args, deleteOK, updateOK, navigateBack, resetOutletContext, requestBooking]);

  const bookingContent = booking && <BookingForm bkng={booking} dte={newDate} onSave={save} onCancel={navigateBack} onDelete={del} onEditDate={onChooseNewDate} />;
  const errorModal = error && <ErrorHandler error={error} />;

  return (
    <Offcanvas show={!chooseNewDate} scroll={!chooseNewDate} backdrop={!chooseNewDate} placement="end" onHide={navigateBack}>
      <Offcanvas.Header closeButton>
        <Offcanvas.Title>Edit Booking</Offcanvas.Title>
      </Offcanvas.Header>
      <Offcanvas.Body>

        {errorModal}
        {bookingContent}

      </Offcanvas.Body>
    </Offcanvas>
  );
}

export default Booking;
