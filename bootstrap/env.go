package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBName     string `mapstructure:"DB_NAME"`
	DBPass     string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	ServerHost string `mapstructure:"SERVER_HOST"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
