package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config is config model for app.
type Config struct {
	// HTTP port.
	Port string

	// Database ip and port.
	Address string
	// Database name.
	DB string
	// Database schema name.
	Schema string
	// Database username.
	User string
	// Database password.
	Password string
}

const (
	// DefaultPort is default HTTP app port.
	DefaultPort = "32001"
	// EnvPath is .env file path.
	EnvPath = "../../config/config.env"
	// EnvPrefix is environment
	EnvPrefix = "TC"
	// DefaultMaxIdleConn is default max db idle connection.
	DefaultMaxIdleConn = 10
	// DefaultMaxOpenConn is default max db open connection.
	DefaultMaxOpenConn = 10
	// DefaultConnMaxLifeTime is default db connection max life time.
	DefaultConnMaxLifeTime = 5 * time.Minute
)

// GetConfig to get config from env.
func GetConfig() (cfg Config) {
	cfg.Port = DefaultPort

	// Load .env file if exist.
	godotenv.Load(EnvPath)

	// Convert env to struct.
	envconfig.Process(EnvPrefix, &cfg)

	// Prepare the ":" for starting HTTP.
	cfg.Port = ":" + cfg.Port

	return cfg
}
