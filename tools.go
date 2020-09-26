package structbot

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

const (
	FolderName = "config"
	FileName   = "config"
)

var configPath string

func init() {
	workDir, _ := os.Getwd()
	configPath = filepath.Join(workDir, FolderName)
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

func makeConfigDirectory() error {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return os.Mkdir(configPath, os.ModePerm)
	}
	return err
}

//保留功能 未來可以新增
func mackConfigFile() error {
	if err := makeConfigDirectory(); err != nil {
		return err
	}
	return nil
}

func readConfigFile() (b []byte, err error) {
	filePath := filepath.Join(configPath, FileName) + ".yaml"
	return ioutil.ReadFile(filePath)
}
