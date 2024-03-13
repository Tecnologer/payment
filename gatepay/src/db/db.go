package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DefaultConnection() (*gorm.DB, error) {
	return NewConnection(DefaultConfig())
}

func NewConnection(config Config) (*gorm.DB, error) {
	dsn := config.String()

	logrus.Debugf("connecting to database: %s", dsn)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
