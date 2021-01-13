package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/router"
)

func main() {

	r := gin.Default()

	// db
	router.StartDatabase(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
