package sysconfigbot

import (
	"errors"
	"reflect"
)

type Bot struct {
	ConfigDir string
}

var ErrInvalidSpecification = errors.New("specification must be a struct pointer")

func checkStruct(spec interface{}) (*reflect.Value, error) {
	s := reflect.ValueOf(spec)
	if s.Kind() != reflect.Ptr {
		return nil, ErrInvalidSpecification
	}
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return nil, ErrInvalidSpecification
	}
	return &s, nil
}

func getTag(elem *reflect.Value) (out []ConfigureFileType) {
	ft := []ConfigureFileType{Yaml, Json, Env, Xml}
	field := elem.Type().Field(0).Tag
	for _, typeId := range ft {
		if _, ok := field.Lookup(typeId.getTagString()); ok {
			out = append(out, typeId)
		}
	}
	return out
}
