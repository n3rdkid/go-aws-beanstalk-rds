package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	DBUsername  string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASS"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
}

var env = Env{}

func GetEnv() Env {
	return env
}

func NewEnv() Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to read configuration file!!!")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Failed to un-marshall environment data :: ", err)
	}

	return env
}
