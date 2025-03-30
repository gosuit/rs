package rs

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Config is type for database connection
type Config struct {
	Host     string `confy:"host" yaml:"host" json:"host" toml:"host" env:"REDIS_HOST"      env-default:"localhost"`
	Port     int    `confy:"port" yaml:"port" json:"port" toml:"port" env:"REDIS_PORT"      env-default:"6379"`
	DBNumber int    `confy:"db"   yaml:"db"   json:"db"   toml:"db"   env:"REDIS_DB_NUMBER" env-default:"0"`
	Password string `confy:"password" env:"REDIS_PASSWORD"`
}

func getConfig(cfg *Config) *redis.Options {
	return &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DBNumber,
	}
}

// ConnectToRedis returns a client to the Redis Server specified by Options.
func New(ctx context.Context, cfg *Config) (Client, error) {
	config := getConfig(cfg)

	client := redis.NewClient(config)

	if err := client.Ping(ctx).Err(); err != nil {
		return Client{}, err
	}

	return Client(*client), nil
}
