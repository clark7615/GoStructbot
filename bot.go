package structbot

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"reflect"

	"gopkg.in/yaml.v3"
)

func FileMakeStruct(filePath string, out interface{}) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = MakeStruct(b, out)
	return err
}

func MakeStruct(src interface{}, out interface{}) error {
	b, err := interface2Bytea(src)
	if err != nil {
		return err
	}
	value, err := checkStruct(out)
	if err != nil {
		return err
	}
	data := validData(b, getTag(value))
	switch data {
	case Yaml:
		if err := yaml.Unmarshal(b, out); err != nil {
			return err
		}
	case Json:
		if err := json.Unmarshal(b, out); err != nil {
			return err
		}
	case Xml:
		if err := xml.Unmarshal(b, out); err != nil {
			return err
		}
	default:
		return errors.New("input data can not be unmarshal,Please confirm the struct and tag")
	}
	return nil
}

func interface2Bytea(src interface{}) (out []byte, err error) {
	s := reflect.ValueOf(src).Kind()
	switch s {
	case reflect.Slice:
		out = src.([]byte)
	case reflect.String:
		out = []byte(src.(string))
	case reflect.Map:
		out, _ = json.Marshal(src)
	default:
		return out, errors.New("input data not []byte or string")
	}
	return out, err
}
