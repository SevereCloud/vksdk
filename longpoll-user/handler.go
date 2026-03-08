package longpoll

// EventNewFunc struct.
type EventNewFunc func([]any) error

// FuncList struct.
//
// Deprecated: FuncList built into LongPoll.
type FuncList map[int][]EventNewFunc

// Handler func.
func (funcList FuncList) Handler(event []any) error {
	key := int(event[0].(float64))

	for _, f := range funcList[key] {
		err := f(event)
		if err != nil {
			return err
		}
	}

	return nil
}
