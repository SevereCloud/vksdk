package longpoll

// EventNewFunc struct
type EventNewFunc func([]interface{})

// FuncList struct
type FuncList map[int][]EventNewFunc

// Handler func
func (funcList FuncList) Handler(event []interface{}) {
	key := int(event[0].(float64))
	for _, f := range funcList[key] {
		f(event)
	}
}
