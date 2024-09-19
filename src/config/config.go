package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type ServerConfigs struct {
	ServerAddress string
}

func GetServerConfigs() ServerConfigs {
	config := ServerConfigs{}

	// config.Secret = os.Getenv("SECRET")
	// if config.Secret == "" {
	// 	newSecret, err := utils.RandomString(32)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	config.Secret = newSecret
	// 	fmt.Println("Generated new SECRET:", config.Secret)
	// }

	envPort := os.Getenv("PORT")
	if envPort != "" {
		config.ServerAddress = ":" + envPort
	} else {
		config.ServerAddress = ":3000"
	}

	return config
}
