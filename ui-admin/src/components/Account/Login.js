import { useState, useEffect } from 'react';

import { useNavigate, useLocation } from 'react-router-dom';

import service from '../../hooks/service';
import { setToken } from '../../utils/token';
import ErrorHandler from '../ErrorHandler';

import '../../assets/style/login.css';

function Login() {
  const [token, storeToken] = useState();
  const [rememberMe, setRememberMe] = useState(false);

  const location = useLocation();
  const navigate = useNavigate();
  const pathname = location.state?.from?.pathname;

  const { data, error, loading, request: requestToken } = service.auth.useLogin();

  useEffect(() => {
    if (!loading && data) {
      setToken(data.token, rememberMe);
      storeToken(data.token);
    }

    if (token) {
      const from = pathname || '/';
      navigate(from, { replace: true });
    }
  }, [data, loading, token, pathname, navigate, rememberMe]);

  const errorStatus401 = error && error.status === 401;
  const errorModal = error && !errorStatus401 && <ErrorHandler error={error} />;

  const submit = (e) => {
    e.preventDefault();

    const {
      username: { value: username },
      password: { value: password },
      rememberMe: { checked: rememberMeValue },
    } = e.target.elements;
    setRememberMe(rememberMeValue);

    const args = {
      body: { username, password },
    };
    requestToken(args);
  };

  return (
    <>
      {errorModal}

      <form className="form-login text-center" data-testid="loginform" onSubmit={submit}>
        <h1 className="h3 mb-3 fw-normal">Login</h1>

        <div className="form-floating">
          <input type="username" className="form-control" name="username" id="username" placeholder="Username" required />
          <label htmlFor="username">Username</label>
        </div>
        <div className="form-floating">
          <input type="password" className="form-control" name="password" id="password" placeholder="Password" required />
          <label htmlFor="password">Password</label>
        </div>
        <div className="checkbox mb-3">
          <label className="d-flex justify-content-start">
            <input type="checkbox" className="mx-1 mt-1" value="remember-me" name="rememberMe" />Remember me
          </label>
        </div>
        {errorStatus401 && (
          <div className="mb-2">
            <center>
              <small className="text-danger">Invalid username and/or password.</small>
            </center>
          </div>
        )}
        <button className="w-100 btn btn-lg btn-dark" type="submit">Login</button>
      </form>

    </>
  );
}

export default Login;
