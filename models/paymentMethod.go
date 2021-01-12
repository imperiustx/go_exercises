package models

import "gorm.io/gorm"

// PaymentMethod payment
type PaymentMethod struct {
	gorm.Model
	BankName string `json:"bank_name"`
	UserName string `json:"user_name"`
	Number   string `json:"number"`
}
