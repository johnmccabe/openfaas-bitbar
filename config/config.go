package config

import (
	"log"
	"os"
	"path"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// DefaultDir TODO
const DefaultDir string = "~/.openfaas"
const yamlFile = "config.yaml"
const legacyTOMLFile = "config.toml"

// DefaultStack TODO
var DefaultStack = Stack{
	Name:       "My OpenFaaS",
	Gateway:    "http://localhost:8080",
	Prometheus: "http://localhost:9090",
}

// Dir TODO
func Dir() string {
	cfgPath, _ := homedir.Expand(DefaultDir)
	return path.Clean(cfgPath)
}

// File TODO
func File() string {
	return path.Clean(filepath.Join(Dir(), yamlFile))
}

// LegacyTOMLExists TODO
func LegacyTOMLExists() bool {
	tomlFile := path.Clean(filepath.Join(Dir(), legacyTOMLFile))
	if _, err := os.Stat(tomlFile); !os.IsNotExist(err) {
		return true
	}
	return false
}

// EnsureConfigDir creates a configDir() if it doesn't already exist
func EnsureConfigDir() error {
	dir := Dir()
	if stat, err := os.Stat(dir); err == nil && stat.IsDir() {
		return nil
	}
	err := os.Mkdir(dir, 0700)
	if err != nil {
		return err
	}
	return nil
}

// Read config from the specified dir returning a slice of OpenFaaS instances
func Read() (Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(Dir())

	err := viper.ReadInConfig()
	if err != nil {
		return Config{
			Stacks: []Stack{DefaultStack},
		}, nil
	}

	conf := new(Config)

	err = mapstructure.Decode(viper.AllSettings(), conf)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if len(conf.Stacks) == 0 {
		return Config{
			Stacks: []Stack{DefaultStack},
		}, nil
	}

	return *conf, nil
}
