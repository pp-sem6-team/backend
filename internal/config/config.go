package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string

	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func Load() *Config {
	return &Config{
		Postgres: PostgresConfig{
			Host:     getEnv("POSTGRES_HOST"),
			Port:     getEnv("POSTGRES_PORT"),
			User:     getEnv("POSTGRES_USER"),
			Password: getEnv("POSTGRES_PASSWORD"),
			DBName:   getEnv("POSTGRES_DB"),
			SSLMode:  getEnv("POSTGRES_SSLMODE"),
			TimeZone: getEnv("POSTGRES_TIMEZONE"),

			MaxOpenConns:    getEnvInt("POSTGRES_MAX_OPEN_CONNS"),
			MaxIdleConns:    getEnvInt("POSTGRES_MAX_IDLE_CONNS"),
			ConnMaxLifetime: time.Duration(getEnvInt("POSTGRES_CONN_MAX_LIFETIME_MINUTES")) * time.Minute,
		},
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("env variable not set: " + key)
	}
	return value
}

func getEnvInt(key string) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("env variable not set: " + key)
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		panic("invalid int value for env " + key)
	}

	return v
}
