package query

import (
	"encoding/json"
	"testing"
)

func TestMedia_FilterByID(t *testing.T) {
	a := NewMedia()
	a.FilterByID(113425)

	data, err := json.MarshalIndent(a, " ", "\t")
	if err != nil {
		t.Error(err)
	}

	t.Log(string(data))
}
