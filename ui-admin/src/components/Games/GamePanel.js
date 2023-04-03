import { useCallback, useEffect, useState } from 'react';

import { Tab, Tabs } from 'react-bootstrap';
import { useNavigate, useParams } from 'react-router-dom';

import service from '../../hooks/service';
import Calendar from '../Calendar/Calendar';
import ErrorHandler from '../ErrorHandler';
import Schedule from '../Schedule/Schedule';
import Toast from '../UI/Toast';
import GameForm from './GameForm';

function GamePanel() {
  const params = useParams();
  const { gid } = params;

  const [activeTab, setActiveTab] = useState('1');
  const handleSelectTab = (tab) => setActiveTab(tab);

  const navigate = useNavigate();
  const navigateBack = useCallback(() => navigate('/games', { replace: true }), [navigate]);

  const { data: game = {}, error, request: requestGame } = service.games.useByID();

  useEffect(() => {
    const args = { params: { gid } };
    requestGame(args);
  }, [gid, requestGame]);

  const errorModal = error && <ErrorHandler error={error} />;

  return (
    <>
      {errorModal}

      <div className="d-flex align-items-start mt-1 mb-2 border-bottom">
        <button className="btn btn-white" type="button" onClick={navigateBack}>
          <i className="bi bi-arrow-left" />
        </button>

        <h3 className="pt-1 mx-1">{game.name}</h3>
      </div>

      <Tabs id="game-tab" className="mt-1 mb-2 game-tab" activeKey={activeTab} onSelect={handleSelectTab}>

        <Tab id="1" eventKey="1" title="Calendar">
          {activeTab === '1' && (<Calendar gid={gid} />)}
        </Tab>
        <Tab id="2" eventKey="2" title="Schedule">
          {activeTab === '2' && (<Schedule gid={gid} dur={game.duration} />)}
        </Tab>
        <Tab id="3" eventKey="3" title="Details">
          {activeTab === '3' && (<GameDetails gid={gid} game={game} />)}
        </Tab>

      </Tabs>

    </>
  );
}

function GameDetails({ gid, game }) {
  const [key, setKey] = useState(0);
  const handleOnCancel = () => setKey((currKey) => currKey + 1);

  const { ok, error, request: editGame } = service.games.useUpdate();

  const errorModal = error && <ErrorHandler error={error} />;

  const save = (g) => {
    const args = {
      params: { gid },
      body: g,
    };
    editGame(args);
  };

  return (
    <div className="p-2">
      <GameForm key={key} game={game} onSave={save} onCancel={handleOnCancel} />
      {ok && <Toast />}
      {errorModal}
    </div>
  );
}

export default GamePanel;
