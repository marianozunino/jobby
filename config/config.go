package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DatabaseURL is the URL of the database
var DatabaseURL string

// Port is the port of the server
var Port string

// JwtAccessTokenSecret is the secret used to sign the JWT access token
var JwtAccessTokenSecret string

// EntDebug is the debug mode for ent
var EntDebug bool

// InstrospectionEnabled is the introspection mode for gqlgen
var InstrospectionEnabled bool

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

	if os.Getenv("INSTROSPECTION_ENABLED") == "" {
		os.Setenv("INSTROSPECTION_ENABLED", "false")
	}

	if os.Getenv("DATABASE_URL") == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	if os.Getenv("JWT_ACCESS_TOKEN_SECRET") == "" {
		panic("JWT_ACCESS_TOKEN_SECRET environment variable is not set")
	}

	DatabaseURL = os.Getenv("DATABASE_URL")
	Port = os.Getenv("PORT")
	JwtAccessTokenSecret = os.Getenv("JWT_ACCESS_TOKEN_SECRET")
	EntDebug = os.Getenv("ENT_DEBUG") == "true"
	InstrospectionEnabled = os.Getenv("INSTROSPECTION_ENABLED") == "true"
}
