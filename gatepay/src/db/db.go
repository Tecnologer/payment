package db

import (
	"time"

	gormlogruslogger "github.com/aklinkert/go-gorm-logrus-logger"
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

	logger := gormlogruslogger.NewGormLogrusLogger(
		logrus.StandardLogger().WithField("component", "gorm"),
		100*time.Millisecond,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
}
