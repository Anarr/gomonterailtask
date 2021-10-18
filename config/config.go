package config

import (
	"github.com/spf13/viper"
	"os"
)

var config *viper.Viper

//Init initalize project configuration
func Init(env string) (*viper.Viper, error)  {
	currentDir, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	conf := viper.New()
	conf.SetConfigType("yaml")
	conf.SetConfigName(env)
	conf.AddConfigPath(currentDir)

	if err := conf.ReadInConfig(); err != nil {
		return nil, err
	}

	return conf, nil
}
