package array

import (
	"reflect"
	"testing"
)

var Array = []Item{
	{Name: "test"},
	{Name: "testing"},
}

func TestFilter(t *testing.T) {
	result := Filter(Array, func(v Item) bool {
		return v.Name == "test"
	})
	if reflect.DeepEqual(result, []Item{{Name: "test"}}) == false {
		t.Error("Filter failed, expected [{test}], got ", result)
	}
}
