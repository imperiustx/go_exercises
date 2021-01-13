package paymentmethodmodel

import (
	"gorm.io/gorm"
)

// PaymentMethod payment
type PaymentMethod struct {
	gorm.Model
	BankName string `json:"bank_name"`
	UserName string `json:"paymentmethod_name"`
	Number   string `json:"number"`
}

// TableName table name
func (PaymentMethod) TableName() string {
	return "payment_methods"
}
