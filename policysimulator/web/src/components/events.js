let events = {}, empty = [];

export function On(type, func, ctx) {
  (events[type] = events[type] || []).push([func, ctx])
}

export function Off(type, func) {
  type || (events = {})
  let list = events[type] || empty,
  i = list.length = func ? list.length : 0
  while(i--) func == list[i][0] && list.splice(i,1)
}

export function Emit(type) {
  var args = empty.slice.call(arguments, 1),
  list = events[type] || empty, i=0, j
  while(j=list[i++]) j[0].apply(j[1], args)
}

export default {
	On: On,
	Off: Off,
	Emit: Emit
}