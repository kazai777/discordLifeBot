package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func LoadConfig() (string, string) {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    return os.Getenv("DISCORD_TOKEN"), os.Getenv("DISCORD_CHANNEL_ID")
}
