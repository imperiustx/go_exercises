package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/middleware"
	"github.com/imperiustx/go_excercises/module/cart/carttransport/gincart"
	"github.com/imperiustx/go_excercises/module/category/categorytransport/gincategory"
	"github.com/imperiustx/go_excercises/module/city/citytransport/gincity"
	"github.com/imperiustx/go_excercises/module/food/foodtransport/ginfood"
	"github.com/imperiustx/go_excercises/module/foodlike/foodliketransport/ginfoodlike"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingtransport/ginfoodrating"
	"github.com/imperiustx/go_excercises/module/image/imagetransport/ginimage"
	"github.com/imperiustx/go_excercises/module/order/ordertransport/ginorder"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailtransport/ginorderdetail"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingtransport/ginordertracking"
	"github.com/imperiustx/go_excercises/module/restaurant/restauranttransport/ginrestaurant"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodtransport/ginrestaurantfood"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantliketransport/ginrestaurantlike"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingtransport/ginrestaurantrating"
	"github.com/imperiustx/go_excercises/module/upload/uploadtransport/ginupload"
	"github.com/imperiustx/go_excercises/module/user/usertransport/ginuser"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddresstransport/ginuseraddress"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokentransport/ginuserdevicetoken"
)

func setupRouter(r *gin.Engine, appCtx appctx.AppContext) {
	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))

	v1.POST("/upload", ginupload.Upload(appCtx))
	// v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	users := v1.Group("/users", middleware.RequiredAuth(appCtx))
	{
		users.GET("", ginuser.List(appCtx))
		users.GET("/:user-id", ginuser.GetProfile(appCtx))
		users.PUT("/:user-id", ginuser.Update(appCtx))
		users.DELETE("/:user-id", ginuser.Delete(appCtx))
	}

	userdevicetokens := v1.Group("/user-device-tokens")
	{
		userdevicetokens.POST("", ginuserdevicetoken.Create(appCtx))
		userdevicetokens.GET("", ginuserdevicetoken.List(appCtx))
		userdevicetokens.GET("/:udt-id", ginuserdevicetoken.Get(appCtx))
		userdevicetokens.PUT("/:udt-id", ginuserdevicetoken.Update(appCtx))
		userdevicetokens.DELETE("/:udt-id", ginuserdevicetoken.Delete(appCtx))
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
		restaurantratings.GET("/:rr-id", ginrestaurantrating.GetRestaurantRating(appCtx))
		restaurantratings.PUT("/:rr-id", ginrestaurantrating.UpdateRestaurantRating(appCtx))
		restaurantratings.DELETE("/:rr-id", ginrestaurantrating.DeleteRestaurantRating(appCtx))
	}

	restaurantlikes := v1.Group("/res-likes")
	{
		restaurantlikes.POST("", ginrestaurantlike.CreateRestaurantLike(appCtx))
		restaurantlikes.GET("", ginrestaurantlike.ListRestaurantLike(appCtx))
		restaurantlikes.GET("/:uid/:rid", ginrestaurantlike.GetRestaurantLike(appCtx))
	}

	restaurantfoods := v1.Group("/res-foods")
	{
		restaurantfoods.POST("", ginrestaurantfood.CreateRestaurantFood(appCtx))
		restaurantfoods.GET("", ginrestaurantfood.ListRestaurantFood(appCtx))
		restaurantfoods.GET("/:rid/:fid", ginrestaurantfood.GetRestaurantFood(appCtx))
		restaurantfoods.DELETE("/:rid/:fid", ginrestaurantfood.DeleteRestaurantFood(appCtx))
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

	useraddresses := v1.Group("/user-addresses")
	{
		useraddresses.POST("", ginuseraddress.CreateUserAddress(appCtx))
		useraddresses.GET("", ginuseraddress.ListUserAddress(appCtx))
		useraddresses.GET("/:ua-id", ginuseraddress.GetUserAddress(appCtx))
		useraddresses.PUT("/:ua-id", ginuseraddress.UpdateUserAddress(appCtx))
		useraddresses.DELETE("/:ua-id", ginuseraddress.DeleteUserAddress(appCtx))
	}

	categories := v1.Group("/categories")
	{
		categories.POST("", gincategory.CreateCategory(appCtx))
		categories.GET("", gincategory.ListCategory(appCtx))
		categories.GET("/:cat-id", gincategory.GetCategory(appCtx))
		categories.PUT("/:cat-id", gincategory.UpdateCategory(appCtx))
		categories.DELETE("/:cat-id", gincategory.DeleteCategory(appCtx))
		categories.GET("/:cat-id/foods", gincategory.ListFood(appCtx))
	}

	orders := v1.Group("/orders")
	{
		orders.POST("", ginorder.CreateOrder(appCtx))
		orders.GET("", ginorder.ListOrder(appCtx))
		orders.GET("/:order-id", ginorder.GetOrder(appCtx))
		orders.PUT("/:order-id", ginorder.UpdateOrder(appCtx))
		orders.DELETE("/:order-id", ginorder.DeleteOrder(appCtx))
	}

	orderdetails := v1.Group("/order-details")
	{
		orderdetails.POST("", ginorderdetail.CreateOrderDetail(appCtx))
		orderdetails.GET("", ginorderdetail.ListOrderDetail(appCtx))
		orderdetails.GET("/:od-id", ginorderdetail.GetOrderDetail(appCtx))
		orderdetails.PUT("/:od-id", ginorderdetail.UpdateOrderDetail(appCtx))
		orderdetails.DELETE("/:od-id", ginorderdetail.DeleteOrderDetail(appCtx))
	}

	ordertrackings := v1.Group("/order-trackings")
	{
		ordertrackings.POST("", ginordertracking.CreateOrderTracking(appCtx))
		ordertrackings.GET("", ginordertracking.ListOrderTracking(appCtx))
		ordertrackings.GET("/:ot-id", ginordertracking.GetOrderTracking(appCtx))
		ordertrackings.PUT("/:ot-id", ginordertracking.UpdateOrderTracking(appCtx))
		ordertrackings.DELETE("/:ot-id", ginordertracking.DeleteOrderTracking(appCtx))
	}

	images := v1.Group("/images")
	{
		images.POST("", ginimage.CreateImage(appCtx))
		images.GET("/:usr-id", ginimage.GetImage(appCtx))
		images.DELETE("/:usr-id", ginimage.DeleteImage(appCtx))
	}

	carts := v1.Group("/carts")
	{
		carts.POST("", gincart.Create(appCtx))
		carts.GET("", gincart.List(appCtx))
		carts.GET("/:u-id/:f-id", gincart.GetProfile(appCtx))
		carts.PUT("/:u-id/:f-id", gincart.Update(appCtx))
		carts.DELETE("/:u-id/:f-id", gincart.Delete(appCtx))
	}

	r.Static("/upload", "./static") // Use this for upload files in local server
}

func setupAdminRouter(r *gin.Engine, appCtx appctx.AppContext) {
	r.Use(middleware.Recover(appCtx))
	//admin := r.Group("/v1/admin")
}
