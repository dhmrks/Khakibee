import { useEffect, useState } from 'react';

import { Modal, Button } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';

function ErrorHandler({ error }) {
  const [message, setMessage] = useState({});
  const [show, setShow] = useState(true);
  const navigate = useNavigate();

  const handleClose = () => {
    setShow(false);
    navigate(-1);
  };

  useEffect(() => {
    const handleError = () => {
      switch (error.status) {
        case 403:
          setMessage({ title: error.message, icon: 'bi bi-exclamation-triangle-fill', body: 'Unfortunetly you don’t have permission to access this page.' });
          break;
        case 404:
          setMessage({ title: error.message, icon: 'bi bi-exclamation-diamond-fill', body: 'We couldn’t found the page you are looking for..' });
          break;
        case 422:
          setMessage({ title: error.message, icon: 'bi bi-exclamation-triangle-fill', body: 'The server was unable to process the contained instructions.' });
          break;
        case 500:
          setMessage({ title: error.message, icon: 'bi bi-x-square-fill', body: 'The server encountered an unexpected condition that prevented it from fulfilling the request.' });
          break;
        default:
          setMessage({ title: error.message, icon: 'bi bi- exclamation - diamond - fill', body: 'Unexpected error, our technical team work on it.' });
          break;
      }
    };

    handleError();
  }, [error]);

  return (
    <Modal size="md" aria-labelledby="contained-modal-title-vcenter" centered show={show}>
      <Modal.Header closeButton onHide={handleClose}>
        <Modal.Title id="contained-modal-title-vcenter">
          {message.title}
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <div className="row">
          <div className="col-1">
            <i className={message.icon} style={{ fontSize: '25px' }} />
          </div>
          <div className="col-11">
            <h4>{message.body}</h4>
          </div>
        </div>
      </Modal.Body>
      <Modal.Footer>
        <Button onClick={handleClose}>Back</Button>
      </Modal.Footer>
    </Modal>
  );
}

export default ErrorHandler;
