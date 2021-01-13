package paymentmethodstorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

// NewSQLStore is
func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
