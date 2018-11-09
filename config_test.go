package config

import (
	"errors"
	"reflect"
	"testing"
)

type testObject struct {
	Title       string  `json:"title"`
	FloatNumber float32 `json:"float_number"`
	IntNumber   int     `json:"int_number"`
	Swith       bool    `json:"switch"`
}

var testVar = &testObject{
	Title:       "testing",
	FloatNumber: 42.0,
	IntNumber:   42,
	Swith:       true,
}

func TestSaveConfig(t *testing.T) {
	err := NewConfig("test.json").Save(testVar)
	if err != nil {
		t.Error(err)
	}
}
func TestLoadConfig(t *testing.T) {
	loadedVar := &testObject{}
	err := NewConfig("test.json").Load(loadedVar)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(testVar, loadedVar) {
		t.Error(errors.New("Objects is not equal"))
	}
}
