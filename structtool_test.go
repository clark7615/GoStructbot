package structbot

import (
	"reflect"
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


func Test_getTag(t *testing.T) {
	type args struct {
		elem *reflect.Value
	}
	jst,_:=checkStruct(&struct {
		Id int `json:"id"`
	}{})
	yst,_:=checkStruct(&struct {
		Id int `yaml:"id"`
	}{})
	xst,_:=checkStruct(&struct {
		Id int `xml:"id"`
	}{})
	est,_:=checkStruct(&struct {
		Id int `env:"id"`
	}{})
	tal,_:=checkStruct(&struct {
		Id int `env:"id" json:"id" yaml:"id" xml:"id"`
	}{})
	tests := []struct {
		name    string
		args    args
		wantOut []SerializationType
	}{
		{
			name: "JsonTag",
			args: args{
				elem: jst,
			},
			wantOut: []SerializationType{Json},
		},{
			name: "YamlTag",
			args: args{
				elem: yst,
			},
			wantOut: []SerializationType{Yaml},
		},{
			name: "XmlTag",
			args:args{
				elem: xst,
			},
			wantOut: []SerializationType{Xml},
		},{
			name: "EnvTag",
			args:args{
				elem: est,
			},
			wantOut: []SerializationType{Env},
		},{
			name: "AllTag",
			args:args{
				elem: tal,
			},
			wantOut: []SerializationType{Yaml,Json,Xml,Env},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := getTag(tt.args.elem); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("getTag() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}