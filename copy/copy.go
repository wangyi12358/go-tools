package copy

import (
	"fmt"
	"reflect"
)

type Config struct {
}

func Copy(source interface{}, target interface{}, conf *Config) error {
	var (
		targetValue = reflect.Indirect(reflect.ValueOf(target))
		sourceValue = reflect.Indirect(reflect.ValueOf(source))
	)

	if !targetValue.CanAddr() {
		return ErrUnaddressable
	}

	if !sourceValue.IsValid() {
		return ErrInvalidSource
	}
	fmt.Printf("sourceValue.Type(): %s", sourceValue.Kind())
	//sourceType, isPtrFrom := indirectType(sourceValue.Type())
	//toType, _ := indirectType(to.Type())
	//
	//if fromType.Kind() == reflect.Interface {
	//	fromType = reflect.TypeOf(from.Interface())
	//}
	return nil
}

func indirectType(reflectType reflect.Type) (_ reflect.Type, isPtr bool) {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		reflectType = reflectType.Elem()
		isPtr = true
	}
	return reflectType, isPtr
}
