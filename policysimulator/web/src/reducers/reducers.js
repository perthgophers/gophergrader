import { createStore } from 'redux';
import { combineReducers } from 'redux-immutable';

import Immutable from 'immutable'

import SuburbsReducer from './suburbs.js'

export default function configureStore(initialState) {
  let store = createStore(combineReducers({
		suburbs: SuburbsReducer
  }), Immutable.Map(initialState))
  return store
}