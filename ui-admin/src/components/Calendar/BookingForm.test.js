/* eslint-disable no-undef */
import { renderNode, screen } from '../../testEnv';
import BookingForm from './BookingForm';

const bkng1 = {
  date: '2022-09-26 19:00',
  name: 'Dim Rks',
  email: 'dim13@gmail.com',
  mob_num: '+306972782755',
  status: 'booked',
};
const bkng2 = {
  date: '2022-09-28 19:00',
  name: 'Dimitris Rakas',
  email: 'dim16@gmail.com',
  mob_num: '+306972782758',
  status: 'booked',
};

function getDate() {
  return screen.getByLabelText(/date/i);
}
function getName() {
  return screen.getByLabelText(/full name/i);
}
function getEmail() {
  return screen.getByLabelText(/email/i);
}
function getMobile() {
  return screen.getByLabelText(/mobile number/i);
}

describe('BookingForm', () => {
  const trackOnSave = jest.fn();
  const trackOnEditDte = jest.fn();
  const trackOnCancelForm = jest.fn();

  it('should not save invalid form', async () => {
    renderNode(<BookingForm />);

    expect(getName()).toBeInvalid();
    expect(getEmail()).toBeInvalid();
    expect(getMobile()).toBeInvalid();

    expect(screen.getByTestId('bookingform')).toBeInvalid();
  });

  it('should create a new booking', async () => {
    const { user } = renderNode(<BookingForm dte={bkng1.date} onSave={trackOnSave} />);

    await user.type(getName(), bkng1.name);
    await user.type(getEmail(), bkng1.email);
    await user.type(getMobile(), bkng1.mob_num);

    expect(screen.getByTestId('bookingform')).toBeValid();
    await user.click(screen.getByText(/save/i));
    expect(trackOnSave).toBeCalledWith(bkng1);
  });

  it('should edit an existing booking', async () => {
    const { rerender, user } = renderNode(<BookingForm bkng={bkng1} onSave={trackOnSave} onEditDate={trackOnEditDte} />);

    expect(screen.getByTestId('bookingform')).toBeValid();
    await user.click(screen.getByText(/save/i));
    expect(trackOnSave).toBeCalledWith(bkng1);

    await user.click(getDate());
    expect(trackOnEditDte).toBeCalled();
    expect(trackOnEditDte).toBeCalledTimes(1);
    rerender(<BookingForm bkng={bkng1} dte={bkng2.date} onSave={trackOnSave} onEditDate={trackOnEditDte} />);

    await user.clear(getName());
    await user.type(getName(), bkng2.name);
    await user.clear(getEmail());
    await user.type(getEmail(), bkng2.email);
    await user.clear(getMobile());
    await user.type(getMobile(), bkng2.mob_num);

    expect(screen.getByTestId('bookingform')).toBeValid();
    await user.click(screen.getByText(/save/i));

    expect(trackOnSave).toBeCalledWith(bkng1);
  });

  it('should cancel the booking form', async () => {
    const { user } = renderNode(<BookingForm onCancel={trackOnCancelForm} />);

    await user.click(screen.getByText(/cancel/i));
    expect(trackOnCancelForm).toBeCalled();
    expect(trackOnCancelForm).toBeCalledTimes(1);
  });
});
