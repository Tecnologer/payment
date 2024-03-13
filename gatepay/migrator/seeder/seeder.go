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
		tx := db.Save(customer)
		if tx.Error != nil {
			return errors.Wrapf(tx.Error, "seeder.customers: saving customer %s", customer.Person.Email)
		}
	}

	return nil
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
		tx = db.Save(merchant)
		if tx.Error != nil {
			return errors.Wrapf(tx.Error, "seeder.merchants: saving merchant %s", merchant.Name)
		}

		for _, user := range merchant.Users {
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
