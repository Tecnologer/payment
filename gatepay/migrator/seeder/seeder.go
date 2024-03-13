package seeder

import (
	"deuna.com/payment/gatepay/src/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	seeders := []func(*gorm.DB) error{
		customers,
		merchants,
	}

	db = db.Begin()

	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			_ = db.Rollback()
			return err
		}
	}

	_ = db.Commit()

	return nil
}

func customers(db *gorm.DB) error {
	customersData := []*models.Customer{
		{
			Person: &models.Person{
				Name:  "John Doe",
				Email: "jdoe@outlook.com",
			},
		},
		{
			Person: &models.Person{
				Name:  "John Nommensen",
				Email: "jnommensen@gmail.com",
			},
		},
	}

	for _, customer := range customersData {
		if existsCustomer(db, customer.Person.Email) {
			continue
		}

		tx := db.Save(customer)
		if tx.Error != nil {
			return errors.Wrapf(tx.Error, "seeder.customers: saving customer %s", customer.Person.Email)
		}
	}

	return nil
}

func existsCustomer(db *gorm.DB, email string) bool {
	var customer *models.Customer
	tx := db.Where("email = ?", email).First(&customer)
	return tx.Error == nil
}

func merchants(db *gorm.DB) error {
	merchantsData := []*models.Merchant{
		{
			Name: "Nexus Innovate",
			Users: []*models.MerchantUser{
				{
					Person: &models.Person{
						Name:  "Alex Mercer",
						Email: "alex.mercer@nexusinnovate.com",
					},
					Role: models.Manager,
				},
				{
					Person: &models.Person{
						Name:  "Jordan Lee",
						Email: "jordan.lee@nexusinnovate.com",
					},
					Role: models.Staff,
				},
				{
					Person: &models.Person{
						Name:  "Taylor Smith",
						Email: "taylor.smith@nexusinnovate.com",
					},
					Role: models.SuperAdmin,
				},
			},
		},
	}

	var tx *gorm.DB

	for _, merchant := range merchantsData {
		if existsMerchant(db, merchant.Name) {
			continue
		}

		tx = db.Save(merchant)
		if tx.Error != nil {
			return errors.Wrapf(tx.Error, "seeder.merchants: saving merchant %s", merchant.Name)
		}

		for _, user := range merchant.Users {
			if existsMerchantUser(db, user.Person.Email) {
				continue
			}

			user.MerchantID = merchant.ID

			tx = db.Save(user)
			if tx.Error != nil {
				return errors.Wrapf(
					tx.Error,
					"seeder.merchants: saving user %s for merchant %s",
					user.Email,
					merchant.Name,
				)
			}
		}
	}

	return nil
}

func existsMerchant(db *gorm.DB, name string) bool {
	var merchant *models.Merchant
	tx := db.Where("name = ?", name).First(&merchant)
	return tx.Error == nil
}

func existsMerchantUser(db *gorm.DB, email string) bool {
	var user *models.MerchantUser
	tx := db.Where("email = ?", email).First(&user)
	return tx.Error == nil
}
