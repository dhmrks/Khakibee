import { useEffect, useCallback } from 'react';

import { Offcanvas } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';

import service from '../../hooks/service';
import ErrorHandler from '../ErrorHandler';
import GameForm from './GameForm';

function GameCreate() {
  const navigate = useNavigate();
  const navigateBack = useCallback(() => navigate('/games', { replace: true }), [navigate]);

  const { error, ok, request: newGame } = service.games.useCreate();

  useEffect(() => {
    if (ok) {
      navigateBack();
    }
  }, [ok, navigateBack]);

  const save = (g) => {
    const args = { body: g };
    newGame(args);
  };

  const gameContent = <GameForm onSave={save} onCancel={navigateBack} />;
  const errorModal = error && <ErrorHandler error={error} />;

  return (
    <Offcanvas show onHide={navigateBack} placement="end">
      <Offcanvas.Header closeButton>
        <Offcanvas.Title>Create Game</Offcanvas.Title>
      </Offcanvas.Header>
      <Offcanvas.Body>

        {errorModal}
        {gameContent}

      </Offcanvas.Body>
    </Offcanvas>
  );
}

export default GameCreate;
