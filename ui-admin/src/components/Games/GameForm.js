/* eslint-disable camelcase */

function GameForm({ formKey = 0, game = {}, onSave, onCancel }) {
  let statusBox;
  if (game.status) { statusBox = game.status === 'active'; }

  const submitHandler = (e) => {
    e.preventDefault();
    const {
      status: { checked: active },
      name: { value: name },
      descr: { value: descr },
      addr: { value: addr },
      img_url: { value: img_url },
      map_url: { value: map_url },
      players: { value: players },
      duration: { value: dur },
      age_range: { value: age_range },
    } = e.target.elements;
    const status = active ? 'active' : 'inactive';
    const duration = parseInt(dur, 10);

    onSave({ status, name, descr, addr, img_url, map_url, players, duration, age_range });
  };

  return (

    <form key={formKey} data-testid="gameform" onSubmit={submitHandler}>

      <div className="row d-flex justify-content-start">
        <div className="col-md-4">

          <div className="mb-3">
            <div className="form-check form-switch">
              <label htmlFor="gamestatus" className="form-check-label">Active</label>
              <input type="checkbox" id="gamestatus" name="status" className="form-check-input" defaultChecked={statusBox} />
            </div>
          </div>

          <div className="mb-3">
            <label htmlFor="gamename" className="form-label">Name</label>
            <input type="text" id="gamename" name="name" className="form-control" placeholder="Enter game name" defaultValue={game.name} required />
          </div>

          <div className="mb-3">
            <label htmlFor="gamedescr" className="form-label">Description</label>
            <textarea className="form-control" id="gamedescr" name="descr" rows="3" placeholder="Enter some brief about game.." defaultValue={game.descr} required />
          </div>

          <div className="mb-3">
            <label htmlFor="gameaddress" className="form-label">Address</label>
            <input type="text" id="gameaddress" name="addr" className="form-control" placeholder="Enter game address" defaultValue={game.addr} required />
          </div>

          <div className="mb-3">
            <label htmlFor="gameimage" className="form-label">Image URL</label>
            <input type="url" id="gameimage" name="img_url" className="form-control" placeholder="Enter image URL" defaultValue={game.img_url} required />
          </div>

          <div className="mb-3">
            <label htmlFor="gamemap" className="form-label">Map URL</label>
            <input type="url" id="gamemap" name="map_url" className="form-control" placeholder="Enter map pin URL" defaultValue={game.map_url} required />
          </div>

          <div className="mb-3">
            <label htmlFor="gameplayers" className="form-label">PLayers</label>
            <input type="text" id="gameplayers" name="players" className="form-control" placeholder="Enter number of players allowed (ex. 3-7)" defaultValue={game.players} required />
          </div>

          <div className="mb-3">
            <label htmlFor="gameduration" className="form-label">Duration</label>
            <input type="number" min="30" step="10" id="gameduration" name="duration" className="form-control" placeholder="Enter game duration in minutes" defaultValue={game.duration} required />
          </div>

          <div className="mb-3">
            <label htmlFor="gameage" className="form-label">Age Range</label>
            <input type="text" id="gameage" name="age_range" className="form-control" placeholder="Enter allowed age range" defaultValue={game.age_range} required />
          </div>

          <div className="d-grid gap-2 col-4 mx-auto">
            <button type="submit" className="btn btn-primary btn-sm"> Save </button>
            <button type="button" className="btn btn-link btn-sm me-2" onClick={onCancel}> Clear </button>
          </div>

        </div>
      </div>
    </form>
  );
}

export default GameForm;
