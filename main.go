package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/controllers"
	"github.com/imperiustx/go_excercises/models"
)

func main() {

	r := gin.Default()

	// db
	db, err := models.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:usr-id", controllers.GetUser)
	r.PUT("/users/:usr-id", controllers.UpdateUser)

	r.POST("/addresses", controllers.CreateAddress)
	r.GET("/addresses", controllers.GetAllAddresses)
	r.GET("/addresses/:add-id", controllers.GetAddress)
	r.PUT("/addresses/:add-id", controllers.UpdateAddress)

	r.POST("/payment-methods", controllers.CreatePaymentMethod)
	r.GET("/payment-methods", controllers.GetAllPayments)
	r.GET("/payment-methods/:pm-id", controllers.GetPayment)
	r.PUT("/payment-methods/:pm-id", controllers.UpdatePayment)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
