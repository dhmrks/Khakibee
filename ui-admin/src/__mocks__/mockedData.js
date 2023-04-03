const mockToken = { token: '123ewq!@#EWQ' };
const mockError401 = { status: 401, message: 'Unauthorized' };
const mockCredentials = {
  body: {
    username: 'username',
    password: 'password',
  },
};
const mockGames = [
  {
    game_id: 1,
    status: 'active',
    name: 'Mansion',
    descr: 'The first escape house in Europe !!1',
    addr: 'Charokopou 93, Kallithea, 2st Floor',
    img_url: 'https://res.cloudinary.com/hwrkhvisl/image/upload/v1586116498/Paradox%20Project/iizdo9dvtx21dqmvrpwp.jpg?fbclid=IwAR0f2SG_pt8_iFH3liW-xHJFz3jt8PKNq4nUmd_2hVd0XRFkouFzhy5f4WU',
    map_url: 'https://www.google.com/maps/place/Paradox+Project/@37.9600721,23.7055711,17z/data=!3m1!4b1!4m5!3m4!1s0x14a1bcf8e3a8f505:0xd72de356a1596eff!8m2!3d37.9600721!4d23.7077598',
    players: '3-7',
    duration: 180,
    age_range: '15+',
  },
  {
    game_id: 87,
    status: 'inactive',
    name: 'Academy',
    descr: 'The first escape house in Europe edit',
    addr: 'Charokopou 93, Kallithea, 2st Floor',
    img_url: 'https://res.cloudinary.com/hwrkhvisl/image/upload/v1586116498/Paradox%20Project/iizdo9dvtx21dqmvrpwp.jpg?fbclid=IwAR0f2SG_pt8_iFH3liW-xHJFz3jt8PKNq4nUmd_2hVd0XRFkouFzhy5f4WU',
    map_url: 'https://www.google.com/maps/place/Paradox+Project/@37.9600721,23.7055711,17z/data=!3m1!4b1!4m5!3m4!1s0x14a1bcf8e3a8f505:0xd72de356a1596eff!8m2!3d37.9600721!4d23.7077598',
    players: '3-7',
    duration: 180,
    age_range: '15+',
  },
];

export {
  mockGames,
  mockToken,
  mockError401,
  mockCredentials,
};
