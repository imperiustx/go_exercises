package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/appctx/cfg"
	"github.com/imperiustx/go_excercises/appctx/uploadprovider"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Load configuration
	config, err := cfg.LoadConfig(".")
	if err != nil {
		log.Println("Viper cannot load config")
		return err
	}

	s3Domain := fmt.Sprintf("https://%s", config.S3Domain)

	s3Provider := uploadprovider.NewS3Provider(config.S3BucketName, config.S3Region, config.S3APIKey, config.S3Secret, s3Domain)

	// db
	db, err := gorm.Open(mysql.Open(config.DBConnectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	// appCtx
	appCtx := appctx.NewAppContext(db, config.SecretKey, s3Provider)

	// migration
	// if err := migrate(db); err != nil {
	// 	return err
	// }

	r := gin.Default()
	// router
	setupRouter(r, appCtx)

	return r.Run() // listen and serve on 0.0.0.0:8080
}
