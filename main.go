package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/middleware"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
	"github.com/imperiustx/go_excercises/module/address/addresstransport/ginaddress"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
	"github.com/imperiustx/go_excercises/module/restaurant/restauranttransport/ginrestaurant"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
	"github.com/imperiustx/go_excercises/module/user/usertransport/ginuser"
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

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})
	if err != nil {
		return err
	}

	appCtx := appctx.NewAppContext(db)

	r.Use(middleware.Recover(appCtx))

	// migration
	if err := db.AutoMigrate(&usermodel.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&addressmodel.Address{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&restaurantmodel.Restaurant{}); err != nil {
		return err
	}

	// router
	users := r.Group("/users")
	{
		users.POST("", ginuser.CreateUser(appCtx))
		users.GET("", ginuser.GetAllUsers(appCtx))
		users.GET("/:usr-id", ginuser.GetUser(appCtx))
		users.PUT("/:usr-id", ginuser.UpdateUser(appCtx))
		users.DELETE("/:usr-id", ginuser.DeleteUser(appCtx))
	}

	addresses := r.Group("/addresses")
	{
		addresses.POST("", ginaddress.CreateAddress(appCtx))
		addresses.GET("", ginaddress.GetAllAddresses(appCtx))
		addresses.GET("/:add-id", ginaddress.GetAddress(appCtx))
		addresses.PUT("/:add-id", ginaddress.UpdateAddress(appCtx))
	}

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.GetAllRestaurants(appCtx))
		restaurants.GET("/:res-id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PUT("/:res-id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:res-id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run() // listen and serve on 0.0.0.0:8080
}
