package main

import (
	"flag"

	"deuna.com/payment/gatepay/migrator/seeder"

	db2 "deuna.com/payment/gatepay/src/db"
	"deuna.com/payment/gatepay/src/models"

	"github.com/sirupsen/logrus"
)

var (
	dbName = flag.String("db-name", "gatepay", "Database name")
	dbUser = flag.String("db-user", "postgres", "Database user")
	dbPass = flag.String("db-pass", "S3cret*_2024", "Database password")
	dbHost = flag.String("db-host", "localhost", "Database host")
	dbPort = flag.Int("db-port", 5432, "Database port")
)

func main() {
	flag.Parse()

	dbConfig := db2.Config{
		Database: *dbName,
		User:     *dbUser,
		Password: *dbPass,
		Host:     *dbHost,
		Port:     *dbPort,
	}

	cnn, err := db2.NewConnection(dbConfig)
	if err != nil {
		logrus.WithError(err).Fatal("connecting to database")
	}

	err = cnn.AutoMigrate(
		&models.Customer{},
		&models.Item{},
		&models.Merchant{},
		&models.MerchantUser{},
		&models.Payment{},
		&models.PaymentItem{},
		&models.PaymentMethod{},
		&models.ActivityLog{},
	)
	if err != nil {
		logrus.WithError(err).Fatal("migrating models")
	}

	err = seeder.Run(cnn)
	if err != nil {
		logrus.WithError(err).Fatal("running seeders")
	}

	logrus.Info("migrations completed")
}
