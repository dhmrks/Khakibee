/* eslint-disable no-undef */
import { renderNode, screen } from '../../testEnv';
import GameForm from './GameForm';

const game1 = {
  addr: 'game address',
  age_range: '13+',
  descr: 'game description',
  duration: 180,
  img_url: 'https://image.com',
  map_url: 'https://map.com',
  name: 'game name',
  players: '5 players',
  status: 'active',
};
const game2 = {
  addr: 'game address 2',
  age_range: '15+',
  descr: 'game description 2',
  duration: 240,
  img_url: 'https://image2.com',
  map_url: 'https://map2.com',
  name: 'game name 2',
  players: '3 players',
  status: 'inactive',
};

function getActive() {
  return screen.getByLabelText(/active/i);
}
function getName() {
  return screen.getByLabelText(/name/i);
}
function getDescription() {
  return screen.getByLabelText(/description/i);
}
function getAddress() {
  return screen.getByLabelText(/address/i);
}
function getPlayers() {
  return screen.getByLabelText(/players/i);
}
function getDuration() {
  return screen.getByLabelText(/duration/i);
}
function getAge() {
  return screen.getByLabelText(/age range/i);
}
function getImageUrl() {
  return screen.getByLabelText(/image url/i);
}
function getMapUrl() {
  return screen.getByLabelText(/map url/i);
}

describe('GameForm', () => {
  const trackOnSave = jest.fn();

  it('save invalid form', async () => {
    const { user } = renderNode(<GameForm />);

    expect(getName()).toBeInvalid();
    expect(getDescription()).toBeInvalid();
    expect(getAddress()).toBeInvalid();
    expect(getPlayers()).toBeInvalid();
    expect(getDuration()).toBeInvalid();
    expect(getAge()).toBeInvalid();

    await user.type(getImageUrl(), 'invalid.url');
    expect(getImageUrl()).toBeInvalid();

    await user.type(getMapUrl(), 'invalid.url');
    expect(getMapUrl()).toBeInvalid();

    expect(screen.getByTestId('gameform')).toBeInvalid();
  });

  it('save completed form', async () => {
    const { user } = renderNode(<GameForm onSave={trackOnSave} />);

    await user.click(getActive());
    await user.type(getName(), game1.name);
    await user.type(getDescription(), game1.descr);
    await user.type(getAddress(), game1.addr);
    await user.type(getImageUrl(), game1.img_url);
    await user.type(getMapUrl(), game1.map_url);
    await user.type(getPlayers(), game1.players);
    await user.type(getDuration(), game1.duration.toString());
    await user.type(getAge(), game1.age_range);

    expect(screen.getByTestId('gameform')).toBeValid();
    await user.click(screen.getByText(/save/i));
    expect(trackOnSave).toBeCalledWith(game1);
  });

  it('save completed form with provided game', async () => {
    const { user } = renderNode(<GameForm game={game2} onSave={trackOnSave} />);

    expect(screen.getByTestId('gameform')).toBeValid();
    await user.click(screen.getByText(/save/i));
    expect(trackOnSave).toBeCalledWith(game2);

    await user.click(getActive());
    await user.clear(getName());
    await user.type(getName(), game1.name);
    await user.clear(getDescription());
    await user.type(getDescription(), game1.descr);
    await user.clear(getAddress());
    await user.type(getAddress(), game1.addr);
    await user.clear(getImageUrl());
    await user.type(getImageUrl(), game1.img_url);
    await user.clear(getMapUrl());
    await user.type(getMapUrl(), game1.map_url);
    await user.clear(getPlayers());
    await user.type(getPlayers(), game1.players);
    await user.clear(getDuration());
    await user.type(getDuration(), game1.duration.toString());
    await user.clear(getAge());
    await user.type(getAge(), game1.age_range);

    expect(screen.getByTestId('gameform')).toBeValid();
    await user.click(screen.getByText(/save/i));

    expect(trackOnSave).toBeCalledWith(game1);
  });

  it('clear form', async () => {
    const trackOnCancel = jest.fn();
    const { user } = renderNode(<GameForm onCancel={trackOnCancel} />);

    await user.click(screen.getByText(/clear/i));
    expect(trackOnCancel).toBeCalledTimes(1);
  });
});
