package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() (*Config, error) {
	viper.SetConfigName("app")  // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name

	dirs := []string{"src/conf", "/src/conf/", "$HOME/src/conf/"}
	for _, d := range dirs {
		viper.AddConfigPath(d)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found; ignore error if desired")
		} else {
			log.Println(err)
		}
		return nil, err
	}
	var conf Config
	if err := viper.Unmarshal(&conf); err != nil {
		log.Println(err)
		return nil, err
	}
	return &conf, nil
}
