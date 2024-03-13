package models

type Customer struct {
	*Person
	PaymentHistory []*Payment `json:"payment_history" gorm:"foreignKey:CustomerID;"`
}
