package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
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

	err := os.Setenv("DB_CONNECTION_STRING", "usr:secret@tcp(localhost:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	err = os.Setenv("SECRET_KEY", "secret")
	if err != nil {
		return err
	}
	err = os.Setenv("S3_BUCKET_NAME", "fd-image-storage")
	if err != nil {
		return err
	}
	err = os.Setenv("S3_REGION", "ap-southeast-1")
	if err != nil {
		return err
	}
	err = os.Setenv("S3_API_KEY", "AKIAVWDLKJEF4PMY7BPM")
	if err != nil {
		return err
	}
	err = os.Setenv("S3_SECRET", "moyfSfDHIUUtVARLHnqa0Iuq1YYHPvZfD8JTPoBi")
	if err != nil {
		return err
	}
	err = os.Setenv("S3_DOMAIN", "doz7th414a7s3.cloudfront.net")
	if err != nil {
		return err
	}

	dsn := os.Getenv("DB_CONNECTION_STRING")
	secret := os.Getenv("SECRET_KEY")

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_API_KEY")
	s3Secret := os.Getenv("S3_SECRET")
	s3Domain := fmt.Sprintf("https://%s", os.Getenv("S3_DOMAIN"))

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3Secret, s3Domain)

	// db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// ctx
	appCtx := appctx.NewAppContext(db, secret, s3Provider)

	// migration
	// if err := migrate(db); err != nil {
	// 	return err
	// }

	r := gin.Default()
	// router
	setupRouter(r, appCtx)

	return r.Run() // listen and serve on 0.0.0.0:8080
}
