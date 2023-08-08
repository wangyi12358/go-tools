package array

import (
	"reflect"
	"testing"
)

type Item struct {
	Name string
}

func TestMap(t *testing.T) {
	var array = []Item{
		{Name: "test"},
		{Name: "testing"},
	}
	result := Map(array, func(v Item) string {
		return v.Name
	})
	if reflect.DeepEqual(result, []string{"test", "testing"}) == false {
		t.Error("Map failed, expected [test testing], got ", result)
	}
}
