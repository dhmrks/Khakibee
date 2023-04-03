import { useState } from 'react';

import classnames from 'classnames';
import { Toast as RBToast, ToastContainer } from 'react-bootstrap';

const defaultMessages = {
  header: 'Updated successfully',
  body: 'Your inputs have been saved.',
  icon: 'bi bi-check-circle-fill',
  iconColor: 'text-success',
};

function Toast({ messages = defaultMessages }) {
  const [show, setShow] = useState(true);
  const handleClose = () => setShow(false);

  const toastBodyClassnames = classnames(`${messages.icon} ${messages.iconColor}`, 'mx-2');

  return (
    <ToastContainer position="top-end">
      <RBToast bg="dark" show={show} onClose={handleClose} delay={2500} autohide>
        <RBToast.Header closeButton>
          <strong className="me-auto">
            <i className={toastBodyClassnames} />{messages.header}
          </strong>
        </RBToast.Header>
        <RBToast.Body className="text-white">{messages.body}</RBToast.Body>
      </RBToast>
    </ToastContainer>
  );
}

export default Toast;
