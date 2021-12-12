package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var errorUnknownFileExtension = errors.New("unknown file extension")

func Parse(path string, cfg *Config) (err error) {

	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return parseYaml(path, cfg)
	default:
		return errorUnknownFileExtension
	}

}

func parseYaml(path string, cfg *Config) (err error) {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if e := file.Close(); e == nil {
			err = e
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return nil
}
