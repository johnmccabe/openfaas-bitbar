package config

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type OpenFaaSInstance struct {
	Name       string
	Gateway    string
	Prometheus string
}

// Read config from the specified dir returning a slice of OpenFaaS instances
func Read(dir string) ([]OpenFaaSInstance, error) {
	f, err := homedir.Expand(dir)
	if err != nil {
		return nil, err
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(f)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Config file not found")
	}

	conf := []OpenFaaSInstance{}

	allSettings := viper.AllSettings()
	for k := range allSettings {
		conf = append(conf, OpenFaaSInstance{
			Name:       viper.GetString(fmt.Sprintf("%s.name", k)),
			Gateway:    viper.GetString(fmt.Sprintf("%s.gateway", k)),
			Prometheus: viper.GetString(fmt.Sprintf("%s.prometheus", k)),
		})
	}

	return conf, nil
}
