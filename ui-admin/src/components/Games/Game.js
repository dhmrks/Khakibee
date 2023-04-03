import { useEffect, useCallback } from 'react';

import { Offcanvas } from 'react-bootstrap';
import { useNavigate, useParams } from 'react-router-dom';

import service from '../../hooks/service';
import ErrorHandler from '../ErrorHandler';
import GameForm from './GameForm';

function Game() {
  const params = useParams();
  const { gid } = params;

  const navigate = useNavigate();
  const navigateBack = useCallback(() => navigate('/games', { replace: true }), [navigate]);

  const { data: game, error, request: requestGame } = service.games.useByID();
  const { ok, request: editGame } = service.games.useUpdate();

  const save = (g) => {
    const args = {
      params: { gid },
      body: g,
    };
    editGame(args);
  };

  useEffect(() => {
    if (ok) {
      navigateBack();
      return;
    }

    const args = { params: { gid } };
    requestGame(args);
  }, [requestGame, gid, ok, navigateBack]);

  const gameContent = game && <GameForm game={game} onSave={save} onCancel={navigateBack} />;
  const errorModal = error && <ErrorHandler error={error} />;

  return (
    <Offcanvas show placement="end" onHide={navigateBack}>
      <Offcanvas.Header closeButton>
        {game && <Offcanvas.Title>{game.name}</Offcanvas.Title>}
      </Offcanvas.Header>
      <Offcanvas.Body>

        {errorModal}
        {gameContent}

      </Offcanvas.Body>
    </Offcanvas>
  );
}

export default Game;
