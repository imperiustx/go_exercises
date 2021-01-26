package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/middleware"
	"github.com/imperiustx/go_excercises/module/address/addresstransport/ginaddress"
	"github.com/imperiustx/go_excercises/module/category/categorytransport/gincategory"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/catrestransport/gincategoryrestaurant"
	"github.com/imperiustx/go_excercises/module/city/citytransport/gincity"
	"github.com/imperiustx/go_excercises/module/food/foodtransport/ginfood"
	"github.com/imperiustx/go_excercises/module/foodlike/foodliketransport/ginfoodlike"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingtransport/ginfoodrating"
	"github.com/imperiustx/go_excercises/module/image/imagetransport/ginimage"
	"github.com/imperiustx/go_excercises/module/order/ordertransport/ginorder"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailtransport/ginorderdetail"
	"github.com/imperiustx/go_excercises/module/restaurant/restauranttransport/ginrestaurant"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantliketransport/ginrestaurantlike"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingtransport/ginrestaurantrating"
	"github.com/imperiustx/go_excercises/module/user/usertransport/ginuser"
)

func setupRouter(r *gin.Engine, appCtx appctx.AppContext) {
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("v1")

	users := v1.Group("/users")
	{
		users.POST("", ginuser.CreateUser(appCtx))
		users.GET("", ginuser.ListUser(appCtx))
		users.GET("/:user-id", ginuser.GetUser(appCtx))
		users.PUT("/:user-id", ginuser.UpdateUser(appCtx))
		users.DELETE("/:user-id", ginuser.DeleteUser(appCtx))
		users.PUT("/:user-id/reactive", ginuser.ReactiveUser(appCtx))
		users.POST("/register", ginuser.Register(appCtx))
		users.POST("/login", ginuser.Login(appCtx))
	}

	restaurants := v1.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.GET("/:res-id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.PUT("/:res-id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:res-id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	restaurantratings := v1.Group("/res-ratings")
	{
		restaurantratings.POST("", ginrestaurantrating.CreateRestaurantRating(appCtx))
		restaurantratings.GET("", ginrestaurantrating.ListRestaurantRating(appCtx))
		restaurantratings.GET("/:res-id", ginrestaurantrating.GetRestaurantRating(appCtx))
		restaurantratings.PUT("/:res-id", ginrestaurantrating.UpdateRestaurantRating(appCtx))
		restaurantratings.DELETE("/:res-id", ginrestaurantrating.DeleteRestaurantRating(appCtx))
	}

	restaurantlikes := v1.Group("/res-likes")
	{
		restaurantlikes.POST("", ginrestaurantlike.CreateRestaurantLike(appCtx))
		restaurantlikes.GET("", ginrestaurantlike.ListRestaurantLike(appCtx))
		restaurantlikes.GET("/:uid/:rid", ginrestaurantlike.GetRestaurantLike(appCtx))
		restaurantlikes.DELETE("/:uid/:rid", ginrestaurantlike.DeleteRestaurantLike(appCtx))
	}

	foods := v1.Group("/foods")
	{
		foods.POST("", ginfood.CreateFood(appCtx))
		foods.GET("", ginfood.ListFood(appCtx))
		foods.GET("/:food-id", ginfood.GetFood(appCtx))
		foods.PUT("/:food-id", ginfood.UpdateFood(appCtx))
		foods.DELETE("/:food-id", ginfood.DeleteFood(appCtx))
	}

	foodlikes := v1.Group("/food-likes")
	{
		foodlikes.POST("", ginfoodlike.CreateFoodLike(appCtx))
		foodlikes.GET("", ginfoodlike.ListFoodLike(appCtx))
		foodlikes.GET("/:uid/:fid", ginfoodlike.GetFoodLike(appCtx))
		foodlikes.DELETE("/:uid/:fid", ginfoodlike.DeleteFoodLike(appCtx))
	}

	foodratings := v1.Group("/food-ratings")
	{
		foodratings.POST("", ginfoodrating.CreateFoodRating(appCtx))
		foodratings.GET("", ginfoodrating.ListFoodRating(appCtx))
		foodratings.GET("/:fr-id", ginfoodrating.GetFoodRating(appCtx))
		foodratings.PUT("/:fr-id", ginfoodrating.UpdateFoodRating(appCtx))
		foodratings.DELETE("/:fr-id", ginfoodrating.DeleteFoodRating(appCtx))
	}

	cities := v1.Group("/cities")
	{
		cities.POST("", gincity.CreateCity(appCtx))
		cities.GET("", gincity.ListCity(appCtx))
		cities.GET("/:city-id", gincity.GetCity(appCtx))
		cities.PUT("/:city-id", gincity.UpdateCity(appCtx))
		cities.DELETE("/:city-id", gincity.DeleteCity(appCtx))
	}

	addresses := v1.Group("/addresses")
	{
		addresses.POST("", ginaddress.CreateAddress(appCtx))
		addresses.GET("", ginaddress.ListAddress(appCtx))
		addresses.GET("/:add-id", ginaddress.GetAddress(appCtx))
		addresses.PUT("/:add-id", ginaddress.UpdateAddress(appCtx))
		addresses.DELETE("/:add-id", ginaddress.DeleteAddress(appCtx))
	}

	categories := v1.Group("/categories")
	{
		categories.POST("", gincategory.CreateCategory(appCtx))
		categories.GET("", gincategory.ListCategory(appCtx))
		categories.GET("/:cat-id", gincategory.GetCategory(appCtx))
		categories.PUT("/:cat-id", gincategory.UpdateCategory(appCtx))
		categories.DELETE("/:cat-id", gincategory.DeleteCategory(appCtx))
	}

	orders := v1.Group("/orders")
	{
		orders.POST("", ginorder.CreateOrder(appCtx))
		orders.GET("", ginorder.ListOrder(appCtx))
		orders.GET("/:ord-id", ginorder.GetOrder(appCtx))
		orders.PUT("/:ord-id", ginorder.UpdateOrder(appCtx))
		orders.DELETE("/:ord-id", ginorder.DeleteOrder(appCtx))
	}

	orderdetails := v1.Group("/order-details")
	{
		orderdetails.POST("", ginorderdetail.CreateOrderDetail(appCtx))
		orderdetails.GET("", ginorderdetail.ListOrderDetail(appCtx))
		orderdetails.GET("/:ord-id", ginorderdetail.GetOrderDetail(appCtx))
		orderdetails.PUT("/:ord-id", ginorderdetail.UpdateOrderDetail(appCtx))
		orderdetails.DELETE("/:ord-id", ginorderdetail.DeleteOrderDetail(appCtx))
	}

	categoryrestaurants := v1.Group("/cat-res")
	{
		categoryrestaurants.POST("", gincategoryrestaurant.CreateCategoryRestaurant(appCtx))
		categoryrestaurants.GET("", gincategoryrestaurant.ListCategoryRestaurant(appCtx))
		categoryrestaurants.GET("/:cid/:rid", gincategoryrestaurant.GetCategoryRestaurant(appCtx))
		categoryrestaurants.PUT("/:cid/:rid", gincategoryrestaurant.UpdateCategoryRestaurant(appCtx))
		categoryrestaurants.DELETE("/:cid/:rid", gincategoryrestaurant.DeleteCategoryRestaurant(appCtx))
	}

	images := v1.Group("/images")
	{
		images.POST("", ginimage.CreateImage(appCtx))
		images.GET("/:usr-id", ginimage.GetImage(appCtx))
		images.DELETE("/:usr-id", ginimage.DeleteImage(appCtx))
	}
}

func setupAdminRouter(r *gin.Engine, appCtx appctx.AppContext) {
	r.Use(middleware.Recover(appCtx))
	//admin := r.Group("/v1/admin")
}
