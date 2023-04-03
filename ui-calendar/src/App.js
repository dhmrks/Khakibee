import { useState } from 'react';

import BkngForm from './components/BkngForm';
import Calendar from './components/Calendar';
import Finish from './components/Finish';
import Games from './components/Games';
import Stepper from './components/Stepper';

const initGame = {};

const gameStep = 0;
const calendarStep = 1;
const bkngFormStep = 2;
const finishStep = 3;

function App() {
  const [step, setStep] = useState(gameStep);
  const [game, setGame] = useState(initGame);
  const [event, setEvent] = useState();
  const [bkng, setBkng] = useState();

  const selections = {
    step1: game.name,
    step2: event,
  };

  const handleStepClick = (s) => {
    setStep(s);
    if (s === gameStep) {
      setGame(initGame);
    }
    setEvent();
  };

  const handleReset = () => {
    setStep(gameStep);
    setGame(initGame);
    setEvent();
  };

  const handleGameClick = (g) => {
    setGame(g);
    setStep(calendarStep);
  };

  const handleCalendarClick = (e) => {
    setEvent(e);
    setStep(bkngFormStep);
  };

  const handleBookingComplete = (b) => {
    setBkng(b);
    setStep(finishStep);
  };

  return (
    <div className="App">
      <div className="container">

        <main className="main">

          <Stepper
            step={step}
            selections={selections}
            onClick={handleStepClick}
          />

          <div className="row">
            {step === gameStep && <Games onClick={handleGameClick} />}
            {step === calendarStep && <Calendar gameID={game.game_id} onClick={handleCalendarClick} />}
            {step === bkngFormStep && <BkngForm gameID={game.game_id} event={event} onComplete={handleBookingComplete} />}
            {step === finishStep && <Finish game={game} bkng={bkng} onClick={handleReset} />}
          </div>

        </main>

      </div>
    </div>
  );
}

export default App;
