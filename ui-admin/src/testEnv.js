/* eslint-disable import/no-unresolved */
import '@testing-library/jest-dom';
import { render } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { BrowserRouter } from 'react-router-dom';

import App from './App';

export function renderApp(options) {
  return {
    ...render(<App />, { wrapper: BrowserRouter }),
    user: userEvent.setup({ document, ...options }),
  };
}

export function renderNode(jsx, options) {
  return {
    ...render(jsx),
    user: userEvent.setup({ document, ...options }),
  };
}

export { act, cleanup, renderHook, waitFor, screen } from '@testing-library/react';
