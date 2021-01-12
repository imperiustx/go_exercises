package models

import "gorm.io/gorm"

// UserPaymentMethod address
type UserPaymentMethod struct {
	gorm.Model
	UserID          int `json:"user_id"`
	PaymentMethodID int `json:"payment_method_id"`
}
