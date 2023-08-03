package array

import "testing"

func TestFind(t *testing.T) {
	var array = []int{1, 2, 3, 4, 5}
	var result = Find(array, func(v int) bool {
		return v == 3
	})
	if *result != 3 {
		t.Error("Find failed, expected 3, got ", *result)
	}
}
