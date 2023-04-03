/* eslint-disable no-undef */
import { renderApp, screen, waitFor } from './testEnv';

test('renders learn react link', async () => {
  renderApp();
  await waitFor(() => expect(screen.getByText(/username/i)).toBeInTheDocument());
});
