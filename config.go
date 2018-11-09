package config

import (
	"encoding/json"
	"os"
)

// NewConfig - init config
func NewConfig(fileName string) *jsonConfig {
	return &jsonConfig{fileName: fileName}
}

type jsonConfig struct {
	fileName string
}

func (jc *jsonConfig) Save(data interface{}) error {
	configFile, err := os.Create(jc.fileName)
	defer configFile.Close()
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "    ")

	return encoder.Encode(data)
}

func (jc *jsonConfig) Load(data interface{}) error {
	f, err := os.Open(jc.fileName)
	defer f.Close()

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(data)

	return err
}
