import { useCallback, useEffect, useState } from 'react';

import { useNavigate } from 'react-router-dom';

import service from '../../hooks/service';
import ErrorHandler from '../ErrorHandler';
import GameForm from './GameForm';

function GameCreatePanel() {
  const navigate = useNavigate();
  const navigateBack = useCallback(() => navigate('/games', { replace: true }), [navigate]);

  const [key, setKey] = useState(0);
  const handleOnCancel = () => setKey((currKey) => currKey + 1);

  const { error, ok, request: newGame } = service.games.useCreate();
  const save = (g) => {
    const args = { body: g };
    newGame(args);
  };

  const gameContent = <GameForm key={key} onSave={save} onCancel={handleOnCancel} />;
  const errorModal = error && <ErrorHandler error={error} />;

  useEffect(() => {
    if (ok) {
      navigateBack();
    }
  }, [ok, navigateBack]);

  return (
    <>
      {errorModal}

      <div className="d-flex align-items-start py-1 mb-4 border-bottom px-2">
        <button className="btn btn-white" type="button" onClick={navigateBack}>
          <i className="bi bi-arrow-left" />
        </button>
        <h3 className="mx-2">Create Game</h3>
      </div>

      <div className="mx-3 mb-1">
        <h5 className="mb-1">Setup a new game</h5>
        <small>You should complete the form below in order to setup a new game. The schedule of this game will be initiated as an empty schedule by default, you can module game`s schedule over your needs.</small>
      </div>

      <div className="p-3">
        {gameContent}
      </div>

    </>
  );
}

export default GameCreatePanel;
