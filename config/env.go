package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       int

	DbConnectionString string
	DbName             string

	DiscordClientID     string
	DiscordClientSecret string

	GoogleClientID     string
	GoogleClientSecret string

	SessionCookieKey string
	SessionMaxAge    int
	SessionHttpOnly  bool
	SessionSecure    bool
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		PublicHost:         getEnv("PUBLIC_HOST", "localhost"),
		Port:               getEnvAsInt("PORT", 3000),
		DbConnectionString: getEnv("MONGO_DB_URI", ""),
		DbName:             getEnv("MONGO_DB_NAME", "MONKEYCOM_GO_DB"),

		DiscordClientID:     getEnv("DISCORD_CLIENT_ID", ""),
		DiscordClientSecret: getEnv("DISCORD_CLIENT_SECRET", ""),

		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),

		SessionCookieKey: getEnv("SESSION_COOKIE_KEY", "96a4ce97-efbb-41d3-8a89-ab2dffe91c44"),
		SessionMaxAge:    getEnvAsInt("SESSION_MAX_AGE", 3600),
		SessionHttpOnly:  getEnvAsBool("SESSION_HTTP_ONLY", true),
		SessionSecure:    getEnvAsBool("SESSION_SECURE", false),
	}
}

func getEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return i
	}

	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return b
	}
	return fallback
}
