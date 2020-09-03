package structbot

import (
	"errors"
	"reflect"
)

type Bot struct {
	ConfigDir string
}

var InvalidSpecification = errors.New("specification must be a struct pointer")

func checkStruct(spec interface{}) (*reflect.Value, error) {
	s := reflect.ValueOf(spec)
	if s.Kind() != reflect.Ptr {
		return nil, InvalidSpecification
	}
	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return nil, InvalidSpecification
	}
	return &s, nil
}

func getTag(elem *reflect.Value) (out []SerializationType) {
	ft := []SerializationType{Yaml, Json, Xml,Env }
	field := elem.Type().Field(0).Tag
	for _, typeId := range ft {
		if _, ok := field.Lookup(typeId.getTagString()); ok {
			out = append(out, typeId)
		}
	}
	return out
}
