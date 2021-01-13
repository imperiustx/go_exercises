package router

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	uri = "usr:secret@tcp(localhost:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
)

func connectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
