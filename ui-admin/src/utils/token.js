const setToken = (token, toLocalStorage) => {
  if (toLocalStorage) {
    localStorage.setItem('token', token);
  }

  sessionStorage.setItem('token', token);
};

const getToken = () => {
  if (localStorage.token && !sessionStorage.token) {
    sessionStorage.setItem('token', localStorage.token);
  }

  return sessionStorage.token || null;
};

const removeToken = () => {
  localStorage.removeItem('token');
  sessionStorage.removeItem('token');
};

const isExpiredJWT = (jwt) => jwt.exp < Date.now() / 1000;

export {
  isExpiredJWT,
  removeToken,
  setToken,
  getToken,
};
