import Immutable, {Map} from 'immutable'

export function newFolder(path) {

}

function getParent(path, files) {
	let parts = path.split("/")
	let node = files

	for (var i = 0; i < parts.length-1; i++) {
		node.get('children').get(parts[i])
	}
}

export default (state = Map(), action) => {
	switch(action.type) {
		case "LOAD:FILES":
			return Immutable.fromJS(JSON.parse(action.files))
		case 'MOVE:FILES':
			return state
		default:
			return state;
	}
}