package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//Conf a shared object accross whole project
var Conf = load()

type configuration struct {
	Host string
	Port string
}

func load() *configuration {
	var configuration configuration
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := viper.Unmarshal(&configuration); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return &configuration
}
