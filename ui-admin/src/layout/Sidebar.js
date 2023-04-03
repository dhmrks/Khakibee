import { Dropdown } from 'react-bootstrap';
import { Link, useLocation, useNavigate } from 'react-router-dom';

import { removeToken } from '../utils/token';

function Sidebar() {
  const { pathname } = useLocation();
  const navigate = useNavigate();

  const logOut = () => {
    removeToken();
    navigate('/', { replace: true });
  };

  return (
    <div className="col-auto col-md-3 col-xl-2  bg-dark px-0 sidebar">
      <div className="d-flex flex-column align-items-center align-items-sm-start text-white min-vh-100">
        <Link to="/" className="d-flex align-items-center pb-3 pt-2 mb-md-0 me-md-auto ps-0 ps-sm-3  text-white text-decoration-none">
          <span className="fs-5 d-none d-sm-inline">Menu</span><i className="bi bi-image d-xs-inline d-sm-none" />
        </Link>
        <ul className="nav nav-pills flex-column mb-auto w-100" id="menu">
          <li className={`nav-item ${pathname === '/calendar' ? 'active' : ''}`}>
            <Link to="/calendar" className="nav-link px-0 align-middle text-white">
              <i className="bi bi-calendar3" /> <span className="ms-1 d-none d-sm-inline">Calendar</span>
            </Link>
          </li>
          <li className={`nav-item ${pathname === '/games' ? 'active' : ''}`}>
            <Link to="/games" className="nav-link px-0 align-middle text-white">
              <i className="bi-puzzle" /> <span className="ms-1 d-none d-sm-inline">Games</span>
            </Link>
          </li>
        </ul>
        <Dropdown className="pb-4" drop="end">
          <Dropdown.Toggle variant="link" className="d-flex align-items-center text-white text-decoration-none dropdown-toggle">
            <i className="bi bi-person-circle" />
            <span className="d-none d-sm-inline mx-1">User</span>
          </Dropdown.Toggle>
          <Dropdown.Menu className="dropdown-menu dropdown-menu-dark text-small ">
            <Dropdown.Item>Settings</Dropdown.Item>
            <Dropdown.Item>Profile</Dropdown.Item>
            <Dropdown.Divider />
            <Dropdown.Item onClick={logOut}>Sign out</Dropdown.Item>
          </Dropdown.Menu>
        </Dropdown>
      </div>
    </div>
  );
}

export default Sidebar;
