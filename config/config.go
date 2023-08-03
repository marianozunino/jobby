package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DatabaseURL is the URL of the database
var DatabaseURL string = os.Getenv("DATABASE_URL")

// Port is the port of the server
var Port string = os.Getenv("PORT")

// JwtAccessTokenSecret is the secret used to sign the JWT access token
var JwtAccessTokenSecret string = os.Getenv("JWT_ACCESS_TOKEN_SECRET")

// EntDebug is the debug mode for ent
var EntDebug bool = os.Getenv("ENT_DEBUG") == "true"

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", ":8080")
	}

	if os.Getenv("ENT_DEBUG") == "" {
		os.Setenv("ENT_DEBUG", "false")
	}

	if os.Getenv("DATABASE_URL") == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	if os.Getenv("JWT_ACCESS_TOKEN_SECRET") == "" {
		panic("JWT_ACCESS_TOKEN_SECRET environment variable is not set")
	}
}
