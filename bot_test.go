package sysconfigbot

import (
	"testing"
)

func TestCheckStruct(t *testing.T) {
	type TestStruct struct {
		Id   int    `json:"id"`
		Test string `json:"test"`
	}

	var data = TestStruct{}
	_, err := checkStruct(&data)
	if err != nil {
		t.Errorf("Run test checkStruct function error: %v", err.Error())
	}
	t.Log("Input struct pointer:OK")
	_, err = checkStruct(data)
	if err != nil {
		t.Log("Input struct:OK", "errorMessage:", err.Error())
	}
	_, err = checkStruct(data.Id)
	if err != nil {
		t.Log("Input not a struct:OK", "errorMessage:", err.Error())
	}
	_, err = checkStruct(&data.Id)
	if err != nil {
		t.Log("Input not a struct pointer:OK", "errorMessage:", err.Error())
	}
}

func TestGetTag(t *testing.T) {
	j := struct {
		Id int `env:"id" json:"id" yaml:"id" xml:"id"`
	}{}
	elem, _ := checkStruct(&j)
	value := getTag(elem)
	for _, fileType := range value {
		t.Log(fileType.getTagString())
	}
}
