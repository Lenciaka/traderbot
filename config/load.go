package config

import (
	"github.com/spf13/viper"
)

func LoadConfigFile(path, filename, fileType string) error {
	viper.SetConfigName(filename)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(path)
	return viper.ReadInConfig()
}
