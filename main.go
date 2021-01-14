package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/controllers"
	"github.com/imperiustx/go_excercises/middleware"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
	"github.com/imperiustx/go_excercises/module/address/addresstransport/ginaddress"
	"github.com/imperiustx/go_excercises/module/paymentmethod/paymentmethodmodel"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
	"github.com/imperiustx/go_excercises/module/user/usertransport/ginuser"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
	"github.com/imperiustx/go_excercises/module/userpaymentmethod/userpaymentmethodmodel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	uri = "usr:secret@tcp(localhost:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {

	r := gin.Default()

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	appCtx := appctx.NewAppContext(db)

	r.Use(middleware.Recover(appCtx))

	// migration
	db.AutoMigrate(&usermodel.User{})
	db.AutoMigrate(&useraddressmodel.UserAddress{})
	db.AutoMigrate(&userpaymentmethodmodel.UserPaymentMethod{})
	db.AutoMigrate(&addressmodel.Address{})
	db.AutoMigrate(&paymentmethodmodel.PaymentMethod{})
	db.AutoMigrate(&restaurantmodel.Restaurant{})

	// router
	r.POST("/users", ginuser.CreateUser(appCtx))
	r.GET("/users", ginuser.GetAllUsers(appCtx))
	r.GET("/users/:usr-id", ginuser.GetUser(appCtx))
	r.PUT("/users/:usr-id", ginuser.UpdateUser(appCtx))
	r.DELETE("/users/:usr-id", ginuser.DeleteUser(appCtx))
	r.POST("/users/address", controllers.CreateUserAddress)
	r.POST("/users/payment-method", controllers.CreateUserPaymentMethod)

	r.POST("/addresses", ginaddress.CreateAddress(appCtx))
	r.GET("/addresses", ginaddress.GetAllAddresses(appCtx))
	r.GET("/addresses/:add-id", ginaddress.GetAddress(appCtx))
	r.PUT("/addresses/:add-id", ginaddress.UpdateAddress(appCtx))

	r.POST("/payment-methods", controllers.CreatePaymentMethod)
	r.GET("/payment-methods", controllers.GetAllPayments)
	r.GET("/payment-methods/:pm-id", controllers.GetPayment)
	r.PUT("/payment-methods/:pm-id", controllers.UpdatePayment)

	r.POST("/restaurants", controllers.CreateRestaurant)
	r.GET("/restaurants", controllers.GetAllRestaurants)
	r.GET("/restaurants/:res-id", controllers.GetRestaurant)
	r.PUT("/restaurants/:res-id", controllers.UpdateRestaurant)
	r.DELETE("/restaurants/:res-id", controllers.DeleteRestaurant)

	r.Run() // listen and serve on 0.0.0.0:8080
}
