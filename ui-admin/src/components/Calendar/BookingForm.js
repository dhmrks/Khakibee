/* eslint-disable camelcase */
import { useState } from 'react';

import { isEmptyObj } from '../../utils/misc';
import Modal from '../UI/Modal';

function BookingForm({ bkng = {}, dte, onSave, onCancel, onDelete, onEditDate }) {
  const [showRemoveBkng, setShowRevomeBkng] = useState(false);
  const submitHandler = (e) => {
    e.preventDefault();

    const {
      date: { value: date },
      name: { value: name },
      email: { value: email },
      mob_num: { value: mob_num },
    } = e.target.elements;

    const status = 'booked';
    onSave({ date, status, name, email, mob_num });
  };

  const modalData = {
    header: 'Remove Booking',
    body: <>Are you sure you want to remove the existing booking of<abbr title="bookingName">${bkng.name}</abbr> at <abbr title="bookingDate">${bkng.date}</abbr> ?</>,
    icon: 'bi bi-trash-fill',
    submitAction: 'Remove',
    cancelAction: 'Back',
  };
  const handleShowRemoveBkng = (value) => setShowRevomeBkng(value);
  const removeBkngModal = showRemoveBkng && <Modal data={modalData} onSubmitAction={onDelete} onCancelAction={handleShowRemoveBkng} />;

  return (
    <>
      {removeBkngModal}

      <form data-testid="bookingform" onSubmit={submitHandler}>

        <div className="row">

          <div className="mb-3">
            <label htmlFor="date" className="form-label">Date</label>
            <input type="text" id="date" name="date" className="form-control" placeholder="Select your preference date-time" value={dte || bkng.date} onClick={() => onEditDate(true)} required readOnly />
          </div>

          <div className="mb-3">
            <label htmlFor="name" className="form-label">Full name</label>
            <input type="text" id="name" name="name" className="form-control" placeholder="Enter your full name" defaultValue={bkng.name} required />
          </div>

          <div className="mb-3">
            <label htmlFor="email" className="form-label">Email</label>
            <input type="email" id="email" name="email" className="form-control" placeholder="Enter your email address" defaultValue={bkng.email} required />
          </div>

          <div className="mb-3">
            <label htmlFor="mob_num" className="form-label">Mobile number</label>
            <input type="tel" id="mob_num" name="mob_num" className="form-control" placeholder="Enter your mobile number" defaultValue={bkng.mob_num} required />
          </div>

          <div className="d-flex flex-column">
            {!isEmptyObj(bkng) && <button type="button" className="btn btn-danger btn-sm btn-sm" onClick={() => handleShowRemoveBkng(true)}> Remove </button>}
            <button type="button" className="btn btn-link btn-sm" onClick={onCancel}> Cancel </button>
            <button type="submit" className="btn btn-primary btn-sm"> Save </button>
          </div>

        </div>
      </form>
    </>
  );
}

export default BookingForm;
