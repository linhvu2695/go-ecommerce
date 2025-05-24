package initialization

import (
	"fmt"
	"go-ecommerce/global"
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() {
	env := os.Getenv("APP_ENV")
	if env != "prod" {
		env = "local"
	}
	log.Println("Environment:", env)

	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load configurations: %w", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration: %v", err)
	}
}
