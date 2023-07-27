package config

import (
	"log"
	"os"
	"strconv"

	env "github.com/joho/godotenv"
)

type Config struct {
	PORT  int    `json:"port"`
	DbURL string `json:"db_url"`
}

func GetConfig() *Config {
	if err := env.Load(); err != nil {
		log.Fatal("error while loading env info")
	}
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("error while getting port from .env:", err.Error())
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("connection string is empty!")
	}
	return &Config{
		PORT:  port,
		DbURL: dbUrl,
	}
}
