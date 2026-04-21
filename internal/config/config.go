package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	Postgres PostgresConfig
	Minio    MinioConfig
	JWT      JWTConfig
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

type MinioConfig struct {
	Endpoint     string
	RootUser     string
	RootPassword string
	Bucket       string
	UseSSL       bool
}

type JWTConfig struct {
	Secret          string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
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
		Minio: MinioConfig{
			Endpoint:     getEnv("MINIO_ENDPOINT"),
			RootUser:     getEnv("MINIO_ROOT_USER"),
			RootPassword: getEnv("MINIO_ROOT_PASSWORD"),
			Bucket:       getEnv("MINIO_BUCKET"),
			UseSSL:       getEnvBool("MINIO_USE_SSL"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET"),
			AccessTokenTTL:  time.Duration(getEnvInt("JWT_ACCESS_TTL_SECONDS")) * time.Second,
			RefreshTokenTTL: time.Duration(getEnvInt("JWT_REFRESH_TTL_DAYS")) * 24 * time.Hour,
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

func getEnvBool(key string) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("env variable not set: " + key)
	}

	v, err := strconv.ParseBool(value)
	if err != nil {
		panic("invalid bool value for env " + key)
	}

	return v
}
