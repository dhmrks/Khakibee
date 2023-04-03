import { useEffect } from 'react';

import { useNavigate, useLocation } from 'react-router-dom';

import service from '../../hooks/service';
import ErrorHandler from '../ErrorHandler';

function Games() {
  const navigate = useNavigate();
  const location = useLocation();

  const { data: games, error, request: requestGames } = service.games.useAll();

  useEffect(() => {
    requestGames();
  }, [requestGames, location]);

  const toCreate = () => { navigate('/games/create'); };
  const toGame = (g) => { navigate(`/games/${g.game_id}`); };

  const gameList = games && games.map((g) => <Game key={g.game_id} game={g} setGame={toGame} />);
  const errorModal = error && <ErrorHandler error={error} />;

  return (
    <>
      <h3 className="py-2 mt-1 mb-2 border-bottom">Games</h3>

      <div className="row g-3">

        <div className="col-12">
          <button type="button" className="btn btn-danger" onClick={toCreate}>Create Game</button>
        </div>

        {errorModal}
        {gameList}

      </div>
    </>
  );
}

function Game({ game, setGame }) {
  const onClick = () => setGame(game);
  const onEnterPressed = (e) => e.code === 'Enter' && onClick();

  const statusClass = game.status === 'active' ? 'bg-success' : 'bg-danger';

  return (
    <div className="col">
      <div className="card game" role="button" tabIndex={0} onClick={onClick} onKeyPress={onEnterPressed}>
        <div className="card-body">
          <h6 className="text-truncate">{game.name}</h6>
          <div className={`badge ${statusClass}`}>{game.status}</div>
          <img src={game.img_url} alt={game.name} className="card-img my-3" />
        </div>
      </div>
    </div>
  );
}

export default Games;
