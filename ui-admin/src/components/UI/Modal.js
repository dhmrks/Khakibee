import { Modal as RBModal } from 'react-bootstrap';

function Modal({ data, onSubmitAction, onCancelAction }) {
  return (
    <RBModal size="md" aria-labelledby="contained-modal-title-vcenter" centered show>
      <RBModal.Header className="bg-dark text-white">
        <RBModal.Title id="contained-modal-title-vcenter">
          {data.header}
        </RBModal.Title>
      </RBModal.Header>
      <RBModal.Body>
        <div className="row">
          <div className="col-1">
            <i className={data.icon} style={{ fontSize: '25px' }} />
          </div>
          <div className="col-11">
            {data.body}
          </div>
        </div>
      </RBModal.Body>
      <RBModal.Footer>
        <button type="button" className="btn btn-dark btn-sm me-2" onClick={() => onCancelAction(false)}>{data.cancelAction}</button>
        <button type="button" className="btn btn-danger btn-sm" onClick={onSubmitAction}>{data.submitAction}</button>
      </RBModal.Footer>
    </RBModal>
  );
}

export default Modal;
