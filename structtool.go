package structbot

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"reflect"

	"gopkg.in/yaml.v3"
)

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
	ft := []SerializationType{Json, Yaml, Xml, Env}
	field := elem.Type().Field(0).Tag
	for _, typeId := range ft {
		if _, ok := field.Lookup(typeId.getTagString()); ok {
			out = append(out, typeId)
		}
	}
	return out
}

func validData(data []byte, sType []SerializationType) SerializationType {
	for _, s := range sType {
		switch s {
		case Yaml:
			if err := yaml.Unmarshal(data, &struct{}{}); err != nil {
				continue
			}
			return Yaml
		case Json:
			if json.Valid(data) {
				return Json
			}
			continue
		case Xml:
			if err := xml.Unmarshal(data, &struct{}{}); err != nil {
				continue
			}
			return Xml
		}
	}
	return Unknown
}
