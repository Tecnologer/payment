package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string  `json:"host"     yaml:"host"`
	Port     int     `json:"port"     yaml:"port"`
	User     string  `json:"user"     yaml:"user"`
	Password string  `json:"password" yaml:"password"`
	Database string  `json:"database" yaml:"database"`
	SSLMode  SSLMode `json:"ssl_mode" yaml:"sslMode"`
}

func DefaultConfig() Config {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		logrus.Warnf("DB_PORT is not set, using default port 5432")

		port = 5432
	}

	sslMode, err := SSLModeString(os.Getenv("DB_SSL_MODE"))
	if err != nil {
		logrus.WithError(err).Warnf("using default ssl mode disable")

		sslMode = Disable
	}

	return Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
		SSLMode:  sslMode,
	}
}

func (c Config) String() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Database,
		c.SSLMode,
	)
}
