/* eslint-disable no-undef */
import { BrowserRouter } from 'react-router-dom';

import { mockToken, mockGames, mockCredentials, mockError401 } from '../../__mocks__/mockedData';
import service from '../../hooks/service';
import { renderNode, screen, cleanup } from '../../testEnv';
import * as TokenHelper from '../../utils/token';
import Games from '../Games/Games';
import Login from './Login';

const wrappedLogin = <BrowserRouter><Login /></BrowserRouter>;

function getUserName() {
  return screen.getByLabelText(/username/i);
}
function getPassword() {
  return screen.getByLabelText(/password/i);
}
function getRememberMe() {
  return screen.getByText(/remember me/i);
}
function getLoginButton() {
  return screen.getByRole('button', { name: /login/i });
}

function getGameOne() {
  return screen.getByText(/mansion/i);
}
function getGameTwo() {
  return screen.getByText(/academy/i);
}

const trackSetToken = jest.spyOn(TokenHelper, 'setToken');

jest.mock('../../hooks/service', () => ({
  auth: { useLogin: jest.fn() },
  games: { useAll: jest.fn() },
}));

const mockedNavigator = jest.fn();
const mockUseLocationValue = {
  state: {
    from: {
      pathname: '/games',
    },
  },
};
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockedNavigator,
  useLocation: () => mockUseLocationValue,
}));

const trackRequest = jest.fn();
const initialServiceValues = {
  ok: false,
  data: null,
  error: null,
  loading: false,
  request: trackRequest,
};

describe('Login', () => {
  beforeEach(() => {
    service.auth.useLogin.mockImplementation(() => initialServiceValues);
  });

  afterEach(cleanup);

  it('should render the Login component', async () => {
    renderNode(wrappedLogin);

    expect(screen.getByRole('heading', { name: /login/i })).toBeInTheDocument();
  });

  it('should complete login', async () => {
    const { user } = renderNode(wrappedLogin);

    await user.type(getUserName(), mockCredentials.body.username);
    await user.type(getPassword(), mockCredentials.body.password);
    await user.click(getRememberMe());

    expect(screen.getByTestId('loginform')).toBeValid();
    expect(getRememberMe()).toBeTruthy();
    await user.click(getLoginButton());

    expect(trackRequest).toBeCalledWith(mockCredentials);

    service.auth.useLogin.mockImplementation(() => ({ ...initialServiceValues, ok: true, data: mockToken }));
    renderNode(wrappedLogin);

    expect(trackSetToken).toBeCalledWith(mockToken.token, false);
    expect(mockedNavigator).toHaveBeenCalled();
    expect(mockedNavigator).toHaveBeenCalledWith('/games', { replace: true });
  });

  it('should display an error message on 401 error', async () => {
    const { user } = renderNode(wrappedLogin);

    await user.type(getUserName(), mockCredentials.body.username);
    await user.type(getPassword(), mockCredentials.body.password);
    await user.click(getRememberMe());

    expect(screen.getByTestId('loginform')).toBeValid();
    expect(getRememberMe()).toBeTruthy();
    await user.click(getLoginButton());

    expect(trackRequest).toBeCalledWith(mockCredentials);

    service.auth.useLogin.mockImplementation(() => ({ ...initialServiceValues, ok: false, error: mockError401 }));
    renderNode(wrappedLogin);

    expect(screen.getByText('Invalid username and/or password.')).toBeInTheDocument();
  });
});

describe('A complete flow until rendered all the available games (integration)', () => {
  beforeEach(() => {
    service.auth.useLogin.mockImplementation(() => initialServiceValues);
  });

  it('should complete login and rendered all games', async () => {
    const { user } = renderNode(wrappedLogin);

    await user.type(getUserName(), mockCredentials.body.username);
    await user.type(getPassword(), mockCredentials.body.password);
    await user.click(getRememberMe());

    expect(screen.getByTestId('loginform')).toBeValid();
    expect(getRememberMe()).toBeTruthy();
    await user.click(getLoginButton());

    expect(trackRequest).toBeCalledWith(mockCredentials);

    service.auth.useLogin.mockImplementation(() => ({ ...initialServiceValues, ok: true, data: mockToken }));
    renderNode(wrappedLogin);

    expect(trackSetToken).toBeCalledWith(mockToken.token, false);
    expect(mockedNavigator).toHaveBeenCalled();
    expect(mockedNavigator).toHaveBeenCalledWith('/games', { replace: true });

    expect(trackRequest).toBeCalled();
    service.games.useAll.mockImplementation(() => ({ ...initialServiceValues, ok: true, data: mockGames }));
    renderNode(<Games />);

    expect(getGameOne()).toBeInTheDocument();
    expect(getGameTwo()).toBeInTheDocument();
    expect(mockGames.length).toBe(2);
  });
});
