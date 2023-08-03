package array

import (
	"fmt"
	"testing"
	"time"
)

func TestChunk(t *testing.T) {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var result = Chunk(array, 3)
	var expected = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8}}
	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDate(t *testing.T) {
	now := time.Now()
	now = now.AddDate(0, 1, 0)
	fmt.Println(now.String())
}
