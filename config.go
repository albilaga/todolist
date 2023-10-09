package main

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"log"
)

type databaseConfig struct {
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	Name     string `env:"NAME"`
	SslMode  string `env:"SSL_MODE"`
}

type config struct {
	DatabaseConfig databaseConfig `envPrefix:"DATABASE_"`
}

func getConfig() config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("Error reading .env file", err)
	}
	return cfg
}
