import { Dropdown } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router-dom';

import logo from '../assets/images/reactjs-icon.svg';
import { removeToken } from '../utils/token';

function Navbar() {
  const navigate = useNavigate();

  const logOut = () => {
    removeToken();
    navigate('/', { replace: true });
  };

  return (
    <nav className="navbar navbar-expand-lg bg-dark">
      <div className="container-fluid">

        <Link to="/" className="navbar-brand">
          <img src={logo} alt="logo" width="28" height="22" className="d-inline-block align-text-top" />
        </Link>

        <Dropdown className="d-flex">
          <Dropdown.Toggle variant="link" className="align-items-center text-white text-decoration-none dropdown-toggle">
            <i className="bi bi-person-circle mx-1" />
            <span className="d-none d-sm-inline mx-1">User</span>
          </Dropdown.Toggle>
          <Dropdown.Menu className="dropdown-menu dropdown-menu-dark text-small">
            <Dropdown.Item>Settings</Dropdown.Item>
            <Dropdown.Item>Profile</Dropdown.Item>
            <Dropdown.Divider />
            <Dropdown.Item onClick={logOut}>Sign out</Dropdown.Item>
          </Dropdown.Menu>
        </Dropdown>

      </div>
    </nav>
  );
}

export default Navbar;
