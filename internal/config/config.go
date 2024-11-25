// Package config provides functionality to configure the service
// Currently, only environment variables are supported
package config

import (
	"errors"
	"os"

	"github.com/caarlos0/env/v6"
)

// Config stores server configuration
type Config struct {
	Address               string `env:"MUSIC_ADDRESS" envDefault:"0.0.0.0:8080"`
	SongDetailsAPIAddress string `env:"DETAILS_API_ADDRESS" envDefault:"0.0.0.0:8081"`
	DatabaseStr           string `env:"MUSIC_DATABASE_STRING" envDefault:"postgres://postgres:postgres@host.docker.internal:5432/postgres?sslmode=disable"`
	LogFile               string `env:"MUSIC_LOG_FILE" envDefault:"./log/log.json"`
}

func NewConfig() *Config {
	return &Config{}
}

// GetConfig looks looks for env values for the config
func (c *Config) GetConfig() error {

	env.Parse(c)
	if _, check := os.LookupEnv("MUSIC_ADDRESS"); !check {
		return errors.New("environment variable MUSIC_ADDRESS is not set")
	}
	if _, check := os.LookupEnv("DETAILS_API_ADDRESS"); !check {
		return errors.New("environment variable DETAILS_API_ADDRESS is not set")
	}
	if _, check := os.LookupEnv("MUSIC_DATABASE_STRING"); !check {
		return errors.New("environment variable MUSIC_DATABASE_STRING is not set")
	}
	if _, check := os.LookupEnv("MUSIC_LOG_FILE"); !check {
		return errors.New("environment variable MUSIC_LOG_FILE is not set")
	}
	return nil
}

// GetAddr returns configured address of the service
func (c Config) GetAddr() string {
	return c.Address
}

// getSongDetailsAPIAddress returns address of an API used to get song details
func (c Config) GetSongDetailsAPIAddress() string {
	return c.SongDetailsAPIAddress
}

// getDatabaseConnStr returns postgres connection string
func (c Config) GetDatabaseConnStr() string {
	return c.DatabaseStr
}

// getLogFile returns filepath to the log file
func (c Config) GetLogFile() string {
	return c.LogFile
}
