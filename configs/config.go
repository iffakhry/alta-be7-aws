package configs

import (
	"os"
	"sync"
)

type AppConfig struct {
	Port      string `yaml:"port"`
	SecretJWT string `yaml:"secretjwt"`
	Database  struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig
	defaultConfig.Port = os.Getenv("APP_PORT")
	defaultConfig.SecretJWT = os.Getenv("JWT_SECRET")
	defaultConfig.Database.Driver = os.Getenv("DB_DRIVER")
	defaultConfig.Database.Name = os.Getenv("DB_NAME")
	defaultConfig.Database.Address = os.Getenv("DB_ADDRESS")
	defaultConfig.Database.Port = os.Getenv("DB_PORT")
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")

	// viper.SetConfigType("yaml")
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./configs/")
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Info("failed to open file")
	// 	return &defaultConfig
	// }

	// var finalConfig AppConfig
	// err := viper.Unmarshal(&finalConfig)
	// if err != nil {
	// 	log.Info("failed to extract external config, use default value")
	// 	return &defaultConfig
	// }
	// return &finalConfig
	return &defaultConfig
}
