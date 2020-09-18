package structbot

import (
	"encoding/json"
	"encoding/xml"

	"gopkg.in/yaml.v3"
)

type Bot struct {
}

func (*Bot) MakeStruct(str string, out interface{}) error {
	b := []byte(str)
	value, err := checkStruct(out)
	if err != nil {
		return err
	}
	tag := getTag(value)
	data := validData(b, tag)
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
	}
	return nil
}
