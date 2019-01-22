package config

import (
	"encoding/json"
	"os"
)

// NewConfig - init config
func NewConfig(fileName string) *JSONConfig {
	return &JSONConfig{fileName: fileName}
}

type JSONConfig struct {
	fileName string
}

func (jc *JSONConfig) Save(data interface{}) error {
	configFile, err := os.Create(jc.fileName)
	defer configFile.Close()
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "    ")

	return encoder.Encode(data)
}

func (jc *JSONConfig) Load(data interface{}) error {
	f, err := os.Open(jc.fileName)
	defer f.Close()

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(data)

	return err
}
