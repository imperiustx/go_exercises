package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/middleware"
	"github.com/imperiustx/go_excercises/module/address/addresstransport/ginaddress"
	"github.com/imperiustx/go_excercises/module/restaurant/restauranttransport/ginrestaurant"
	"github.com/imperiustx/go_excercises/module/user/usertransport/ginuser"
)

func setupRouter(r *gin.Engine, appCtx appctx.AppContext) {
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("v1")

	users := v1.Group("/users")
	{
		users.POST("", ginuser.CreateUser(appCtx))
		users.GET("", ginuser.ListUser(appCtx))
		users.GET("/:usr-id", ginuser.GetUser(appCtx))
		users.PUT("/:usr-id", ginuser.UpdateUser(appCtx))
		users.DELETE("/:usr-id", ginuser.DeleteUser(appCtx))
	}

	addresses := v1.Group("/addresses")
	{
		addresses.POST("", ginaddress.CreateAddress(appCtx))
		addresses.GET("", ginaddress.GetAllAddresses(appCtx))
		addresses.GET("/:add-id", ginaddress.GetAddress(appCtx))
		addresses.PUT("/:add-id", ginaddress.UpdateAddress(appCtx))
	}

	restaurants := v1.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.GetAllRestaurants(appCtx))
		restaurants.GET("/:res-id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PUT("/:res-id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:res-id", ginrestaurant.DeleteRestaurant(appCtx))
	}
}

func setupAdminRouter(r *gin.Engine, appCtx appctx.AppContext) {
	r.Use(middleware.Recover(appCtx))
	//admin := r.Group("/v1/admin")
}
