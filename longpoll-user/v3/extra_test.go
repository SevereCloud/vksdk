package wrapper

import (
	"encoding/json"
	"testing"
)

func Test_interfaceToStringIntMap(t *testing.T) {
	t.Parallel()

	var i interface{}

	err := json.Unmarshal([]byte(`{"v":0}`), &i)
	if err != nil {
		t.Error(err)
		return
	}

	m, err := interfaceToStringIntMap(i)
	if err != nil {
		t.Error(err)
		return
	}

	if m["v"] != 0 {
		t.Error("expected v is 0")
	}
}

func Test_interfaceToIDSlice(t *testing.T) {
	t.Parallel()

	var i interface{}

	err := json.Unmarshal([]byte(`[0, 9]`), &i)
	if err != nil {
		t.Error(err)
		return
	}

	m, err := interfaceToIDSlice(i)
	if err != nil {
		t.Error(err)
		return
	}

	if m[1] != 9 {
		t.Error("expected v is 0")
	}
}
