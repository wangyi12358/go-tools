package copy

import (
	"fmt"
	"reflect"
	"testing"
)

type Source struct {
	Name string
}

type Target struct {
	Name string
}

func TestCopy(t *testing.T) {
	//source := &Source{
	//	Name: "source",
	//}
	target := "123"
	tt := &target
	v2 := reflect.ValueOf(tt)
	fmt.Println(v2.Kind())
	Copy(tt, &target, nil)
}
