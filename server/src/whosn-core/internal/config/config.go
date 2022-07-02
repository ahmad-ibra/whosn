package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// Config wraps all the env variables
type Config struct {
	Env            string
	Port           string
	JWTKey         string
	FrontendDomain string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// SvcName returns the service name
const SvcName = "whosn-core"

var lock = &sync.Mutex{}
var cfg *Config

func GetConfig() *Config {
	lock.Lock()
	defer lock.Unlock()
	if cfg == nil {
		cfg = initConfig()
	}
	return cfg
}

// InitConfig returns the whosn-core configuration
// Remember to update .env_example if more env vars have been added
func initConfig() *Config {
	env := os.Getenv("ENV")
	if env != "prod" && env != "test" {
		err := godotenv.Load(".env")
		if err != nil {
			ll := log.WithFields(log.Fields{"function": "InitConfig", "error": err})
			ll.Error("Failed to init config")
			panic(err)
		}
	}

	return &Config{
		Env:            os.Getenv("ENV"),
		Port:           os.Getenv("PORT"),
		JWTKey:         os.Getenv("JWT_KEY"),
		FrontendDomain: os.Getenv("FRONTEND_DOMAIN"),

		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_DBNAME"),
	}
}
