package util

import (
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(filePath string) (config Config, err error) {
	dir := filepath.Dir(filePath)
	file := filepath.Base(filePath)
	fileName := file[:len(file)-len(filepath.Ext(file))]

	viper.AddConfigPath(dir)
	viper.SetConfigName(fileName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)

	return
}
