import React from 'react';

import jwtDecode from 'jwt-decode';
import { Route, Routes, Navigate, useLocation } from 'react-router-dom';

import { isEmptyObj } from './utils/misc';
import { isExpiredJWT, getToken, removeToken } from './utils/token';

// lazy loading page components
const Layout = React.lazy(() => import('./layout/Layout'));
const Login = React.lazy(() => import('./components/Account/Login'));
const Booking = React.lazy(() => import('./components/Calendar/Booking'));
const BookingCreate = React.lazy(() => import('./components/Calendar/BookingCreate'));
const Games = React.lazy(() => import('./components/Games/Games'));
const GameCreate = React.lazy(() => import('./components/Games/GameCreatePanel'));
const GamePanel = React.lazy(() => import('./components/Games/GamePanel'));

const loading = () => <div className="animated fadeIn pt-3 text-center">Loading...</div>;

function App() {
  const booking = (
    <React.Suspense fallback={loading()}>
      <Booking />
    </React.Suspense>
  );
  const bookingCreate = (
    <React.Suspense fallback={loading()}>
      <BookingCreate />
    </React.Suspense>
  );
  const gameCreate = (
    <React.Suspense fallback={loading()}>
      <GameCreate />
    </React.Suspense>
  );
  const gamePanel = (
    <React.Suspense fallback={loading()}>
      <GamePanel />
    </React.Suspense>
  );

  return (
    <div className="App container-fluid">

      <Routes>
        <Route path="/login" element={<Login />} />
        <Route
          path="*"
          element={(
            <PrivateRoute>
              <Layout>
                <React.Suspense fallback={loading()}>

                  <Routes>

                    <Route index path="*" element={<Navigate to="/games" />} />

                    <Route path="/games" element={<Games />} />
                    <Route path="/games/create" element={gameCreate} />

                    <Route path="/games/:gid" element={gamePanel}>
                      <Route path="booking/:bid" element={booking} />
                      <Route path="booking/create" element={bookingCreate} />
                    </Route>

                  </Routes>

                </React.Suspense>
              </Layout>
            </PrivateRoute>
          )}
        />
      </Routes>

    </div>
  );
}

function PrivateRoute({ children }) {
  const location = useLocation();
  const token = getToken();
  const jwt = token ? jwtDecode(token) : {};

  if (isEmptyObj(jwt) || isExpiredJWT(jwt)) {
    removeToken();
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return children;
}

export default App;
