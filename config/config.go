package config

import (
	"log"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// Read config from the specified dir returning a slice of OpenFaaS instances
func Read(dir string) (*Config, error) {
	f, err := homedir.Expand(dir)
	if err != nil {
		return nil, err
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(f)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	conf := new(Config)

	err = mapstructure.Decode(viper.AllSettings(), conf)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return conf, nil
}
