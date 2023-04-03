import { useEffect, useState } from 'react';

import withTranslation from '../HOCs/withTranslation';
import service from '../service';

function BkngForm({ gameID, event, tr: { form }, onComplete }) {
  const [formDtls, setFormDtls] = useState({});
  const { ok, loading, request: createNewBooking } = service.bookings.useCreate();

  useEffect(() => {
    if (ok) { onComplete(formDtls); }
  }, [ok, formDtls, onComplete]);

  const handleSubmit = (e) => {
    e.preventDefault();
    const {
      name: { value: name },
      email: { value: email },
      mobile: { value: mobile },
      notes: { value: notes },
      tncs: { checked: tncs },
    } = e.target.elements;

    const args = {
      params: { gid: gameID },
      body: {
        date: event,
        name,
        email,
        mob_num: mobile,
        notes,
      },
    };
    createNewBooking(args);
    setFormDtls({ name, email, mobile, notes, tncs });
  };

  const loadingCard = (
    <div className="card text-bg-dark opacity-75 position-absolute  h-100 w-100 ">
      <div className="d-flex  align-items-center h-100">
        <div className="row justify-content-center">
          <div className="col-12 text-center">
            <p>{form.process}</p>
          </div>
          <div className="col-12 spinner-border text-light mb-2" role="status" />
          <div className="col-12 text-center">
            <p>{form.loading}</p>
          </div>
        </div>
      </div>
    </div>
  );

  return (
    <div className="col">

      <div className="card">
        { loading && loadingCard}
        <div className="card-body ">
          <form className="row" onSubmit={handleSubmit}>

            <div className="mb-3">
              <label htmlFor="name" className="form-label"><b>{form.name}</b>(*)</label>
              <input type="text" id="name" className="form-control" required />
            </div>

            <div className="mb-3">
              <label htmlFor="email" className="form-label"><b>{form.email}</b>(*)</label>
              <input type="email" id="email" className="form-control" required />
            </div>

            <div className="mb-3">
              <label htmlFor="mobile" className="form-label"><b>{form.mobile}</b>(*)</label>
              <input type="tel" id="mobile" pattern="^\+[1-9]\d{1,14}$" className="form-control" required placeholder="include country code (ex +306975645865)" />
            </div>

            <div className="mb-3">
              <label htmlFor="notes" className="form-label"><b>{form.notes}</b></label>
              <input type="textarea" id="notes" className="form-control" />
            </div>

            <div className="mb-3">
              <div className="form-check">
                <input type="checkbox" id="tncs" name="tncs" className="form-check-input" required />
                <label htmlFor="tncs" className="form-check-label">
                  {form.gdpr1}&nbsp;
                  <a className="bk-title" target="_blank" href="/" rel="noreferrer">
                    {form.tncs}&nbsp;
                  </a>
                  {form.gdpr2}&nbsp;
                  <a className="bk-title" target="_blank" href="/" rel="noreferrer">
                    {form.plc}&nbsp;
                  </a>
                  {form.gdpr3}
                </label>
              </div>
            </div>

            <div>
              <button type="submit" className="btn btn-primary">{form.submitBtn}</button>
            </div>

            <div className="text-muted">
              <small>{form.required}</small>
            </div>

          </form>
        </div>
      </div>
    </div>
  );
}

export default withTranslation(BkngForm);
