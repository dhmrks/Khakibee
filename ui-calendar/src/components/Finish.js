import withTranslation from '../HOCs/withTranslation';

function Finish({ game, bkng, tr: { finish, form }, onClick }) {
  return (
    <div className="col">
      <div className="card">
        <div className="card-header text-center rounded-0">
          <h4>{finish.bookingInfo}</h4>
        </div>
        <div className="card-body">

          <div className="alert alert-success">
            <p className="alert-heading"> <i className="fa fa-check-square icon" /> {finish.checkEmail}</p>
            <p> <i className="fa fa-exclamation-circle icon" /> <strong>{finish.important}</strong>
              {finish.infoAboutMail}
            </p>
          </div>
          <h5>Booking game name and hour
          </h5>
          <a className="fa fa-map-pin bk-title icon clickable" href="google.com" target="_blank" rel="noopener noreferrer">{finish.weAreHere}</a>
          <hr />
          <div>

            {/* <div className="row"> */}
            <div className="col-md-8 mb-3">
              <div className="row">
                <div className="col">{finish.address} </div>
                <div className="col h5"><b>{game.addr}</b></div>
              </div>
              <div className="row">
                <div className="col">{form.name} </div>
                <div className="col"><b>{bkng.name}</b></div>
              </div>
              <div className="row">
                <div className="col">{form.email} </div>
                <div className="col"><b>{bkng.email}</b></div>
              </div>
              <div className="row">
                <div className="col">{form.mobile} </div>
                <div className="col"><b>{bkng.mobile}</b></div>
              </div>
            </div>

            {/* </div> */}
          </div>

          <button type="button" className="btn btn-primary" onClick={onClick}>{finish.newBooking}</button>
        </div>
      </div>
    </div>
  );
}

export default withTranslation(Finish);
