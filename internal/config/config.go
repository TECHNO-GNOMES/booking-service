package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig DatabaseConfig
	Auth     AuthConfig
	S3Config S3Config
}

type DatabaseConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Database string `env:"POSTGRES_DB"`
}
type AuthConfig struct {
	Secret string
}

type S3Config struct {
	S3Endpoint       string `env:"S3_ENDPOINT"`
	S3PublicEndpoint string `env:"S3_PUBLIC_ENDPOINT"`
	S3AccessKey      string `env:"S3_ACCESS_KEY"`
	S3SecretKey      string `env:"S3_SECRET_KEY"`
	S3Bucket         string `env:"S3_BUCKET"`
}

func (cfg DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.Port,
	)
}

func Load() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}

	return &Config{
		DbConfig: DatabaseConfig{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("JWT_SECRET"),
		},
		S3Config: S3Config{
			S3PublicEndpoint: os.Getenv("S3_PUBLIC_ENDPOINT"),
			S3Endpoint:       os.Getenv("S3_ENDPOINT"),
			S3AccessKey:      os.Getenv("S3_ACCESS_KEY"),
			S3SecretKey:      os.Getenv("S3_SECRET_KEY"),
			S3Bucket:         os.Getenv("S3_BUCKET"),
		},
	}
}
