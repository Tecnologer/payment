package models

type Customer struct {
	*Person
	PaymentMethods []*PaymentMethod `json:"payment_methods" gorm:"foreignKey:CustomerID;"`
}
