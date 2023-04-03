function CalendarDay({ day, today, selected, events, onClick }) {
  // if (this.props.date.getTime() === this.props.currentDate.getTime()) { this.props.click(this.props.sessions, this.props.date); }

  let status;
  let disabled;
  if (!events) {
    status = 'disabled';
    disabled = false;
  } else if (selected) {
    status = 'bg-warning';
  } else if (today) {
    status = 'bg-dark';
  }

  const handleClick = () => {
    if (disabled) return;
    onClick(events);
  };

  const handleOnKeyPress = (e) => {
    if (e.code !== 'enter' || disabled) return;
    onClick(events);
  };

  return (
    <div className="weekday-col">
      <div
        role="button"
        tabIndex={day}
        className={`btn btn-success ${status}`}
        disabled={disabled}
        style={{ display: 'block', padding: '0.375rem 0' }}
        onClick={handleClick}
        onKeyPress={handleOnKeyPress}
      >
        {day}
      </div>
    </div>
  );
}

export default CalendarDay;
