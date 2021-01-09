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

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
