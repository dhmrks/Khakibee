import withTranslation from '../HOCs/withTranslation';

import '../assets/style/stepper.css';

const active = 'active';
const completed = 'completed';
const todo = 'todo';
const pencil = 'bi bi-pencil';
const check = 'bi bi-check';
const small = 'small';

function Stepper({ step = 0, selections, onClick, tr: { tabs } }) {
  const step0 = ` ${step === 0 && active} ${step > 0 && completed} ${step < 0 && todo}`;
  const step1 = `${step === 1 && active} ${step > 1 && completed} ${step < 1 && todo}`;
  const step2 = `${step === 2 && active} ${step > 2 && completed} ${step < 2 && todo}`;

  const disableSteps = [0, 1, 2].includes(step);

  const step0Icon = `${small} ${step <= 0 && pencil} ${step > 0 && check}`;
  const step1Icon = `${small} ${step <= 1 && pencil} ${step > 1 && check}`;
  const step2Icon = `${small} ${step <= 2 && pencil} ${step > 2 && check}`;

  return (
    <div className="mt-4 justify-content-center">

      <ul className="stepper stepper-horizontal">

        {/* <!-- Zero Step --> */}
        <li className={step0}>
          <button type="button" className="astep" onClick={() => disableSteps && onClick(0)}>
            <span className="circle">
              <i className={step0Icon} />
            </span>
            <span className="label">{tabs.room}</span>
            <span className="choice"><br />&nbsp;{selections.step1}</span>
          </button>
        </li>

        {/* <!-- First Step --> */}
        <li className={step1}>
          <button type="button" className="astep" onClick={() => disableSteps && onClick(1)}>
            <span className="circle">
              <i className={step1Icon} />
            </span>
            <span className="label">{tabs.date}</span>
            <span className="choice"><br />&nbsp;{selections.step2}</span>
          </button>
        </li>

        {/* <!-- Second Step --> */}
        <li className={step2}>
          <button type="button" className="astep">
            <span className="circle">
              <i className={step2Icon} />
            </span>
            <span className="label">{tabs.form}</span>
            <span className="choice"><br />&nbsp;</span>
          </button>
        </li>
      </ul>
    </div>
  );
}

export default withTranslation(Stepper);
