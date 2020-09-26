package structbot

import (
	"encoding/json"
	"encoding/xml"
	"errors"

	"gopkg.in/yaml.v3"
)

//FileMakeStruct從./config/config.yaml讀取檔案並轉為struct
//此路徑為強迫存在要使用本功能請按照規定建立資料夾以及yaml檔案
func FileMakeStruct(out interface{}) error {
	b, err := readConfigFile()
	if err != nil {
		return err
	}
	err = MakeStruct(b, out)
	return err
}

//MakeStruct將支援的資料型態輸入後可自動賦植給struct
//支援Json、Yaml、Xml輸出的struct必須事先填入相應的TAG才能正常的賦值
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
