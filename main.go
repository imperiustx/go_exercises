package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	uri = "usr:secret@tcp(localhost:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	r := gin.Default()

	// db
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		return err
	}

	// ctx
	appCtx := appctx.NewAppContext(db)

	// migration
	if err := migrate(db); err != nil {
		return err
	}

	// router
	setupRouter(r, appCtx)

	return r.Run() // listen and serve on 0.0.0.0:8080
}
