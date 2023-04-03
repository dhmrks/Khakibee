import { configureStore, combineReducers } from 'redux';

const initialUserState = {
  user_id: null,
};

const userReducer = (action) => {
  let state = initialUserState;
  switch (action.type) {
    case 'SET_USER':
      state = {
        ...state,
        user_id: action.payload,
      };
      break;
    case 'LOG_OUT':
      state = {
        ...state,
        user_id: null,
      };
      break;
    default:
      break;
  }
  return state;
};

export default configureStore(combineReducers({ userReducer }));
