package config

import (
	"fmt"
	"github.com/go-pg/pg"
	"os"
	"strconv"
)

type Config struct {
	Port   string
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrDefaultInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Error parsing %s as integer: %s\n", key, err)
		return defaultValue
	}
	return valueInt
}

func LoadConfigFromEnv() Config {

	var config Config

	config.Port = getEnvOrDefault("PORT", "8080")
	config.DBUser = getEnvOrDefault("DBUSER", "wb")
	config.DBPass = getEnvOrDefault("DBPASS", "wb")
	config.DBHost = getEnvOrDefault("DBHOST", "localhost")
	config.DBPort = getEnvOrDefaultInt("DBPORT", 5434)
	config.DBName = getEnvOrDefault("DBNAME", "wb_l0")

	return config
}

func (cfg *Config) GetDBString() pg.Options {
	return pg.Options{
		User:     cfg.DBUser,
		Password: cfg.DBPass,
		Database: cfg.DBName,
		Addr:     "localhost:5434",
	}
	// return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}
