package longpoll

// EventNewFunc struct.
type EventNewFunc func([]interface{}) error

// FuncList struct.
//
// Deprecated: FuncList built into LongPoll.
type FuncList map[int][]EventNewFunc

// Handler func.
func (funcList FuncList) Handler(event []interface{}) error {
	key := int(event[0].(float64))

	for _, f := range funcList[key] {
		if err := f(event); err != nil {
			return err
		}
	}

	return nil
}
