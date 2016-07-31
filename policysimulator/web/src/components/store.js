import configureStore from '../reducers/reducers'
import Immutable from 'immutable'
import superagent from 'superagent'

let store = configureStore(Immutable.Map())

export default function() {
	return store
}