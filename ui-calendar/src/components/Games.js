import { useEffect } from 'react';

import { Card } from 'react-bootstrap';

import withTranslation from '../HOCs/withTranslation';
import service from '../service';

function Games({ tr: { duration, numOfPlayers }, onClick }) {
  const { data: games, request: requestGames } = service.games.useAll();

  useEffect(() => {
    requestGames();
  }, [requestGames]);

  const gamesList = games && games.map((g) => (
    <div className="col sm-4" key={g.game_id}>
      <Card role="button" onClick={() => onClick(g)}>
        <img src={g.img_url} className="card-img-top photos clickable" alt={g.name} width="100%" />
        <div className="card-body text-center">
          <div className="col lg-12 text-center space-bottom d-grid gap-2">
            <button type="button" className="btn btn-primary">
              <span>{g.name}</span>
            </button>
          </div>
          <div className="d-flex justify-content-between">
            <small><i className="bi bi-hourglass-split" /> <b>{duration} {g.duration}</b></small>
            <small><i className="bi bi-people-fill" /> <b>{numOfPlayers} {g.players}</b></small>
          </div>
        </div>
      </Card>
    </div>
  ));

  // console.log(games);

  return gamesList;
}

export default withTranslation(Games);
