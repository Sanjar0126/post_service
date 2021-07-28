package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment      string // develop, staging, production
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	LogLevel         string
	RPCPort          string
	FareServiceHost  string
	FareServicePort  int
	CronSetting      string
	CDN              string
	MapURL           string
	MapKey           string

	OrderServiceHost string
	OrderServicePort int

	JowiAPIKey    string
	JowiAPISecret string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))

	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "golang1"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "strapi"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "strapi"))

	c.OrderServiceHost = cast.ToString(getOrReturnDefault("ORDER_SERVICE_HOST", "order-service"))
	c.OrderServicePort = cast.ToInt(getOrReturnDefault("ORDER_SERVICE_PORT", 80))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":8003"))
	c.FareServiceHost = cast.ToString(getOrReturnDefault("FARE_SERVICE_HOST", "fare-service"))
	c.FareServicePort = cast.ToInt(getOrReturnDefault("FARE_SERVICE_PORT", 80))
	c.CronSetting = cast.ToString(getOrReturnDefault("CRON_VALUE", "@hourly"))
	c.CDN = cast.ToString(getOrReturnDefault("CDN", "https://cdn.test.delever.uz/delever/"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
