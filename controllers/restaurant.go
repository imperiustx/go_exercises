package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantbusiness"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantstorage"
	"gorm.io/gorm"
)

// CreateRestaurant new restaurant
func CreateRestaurant(c *gin.Context) {
	// TODO: Validate input
	var restaurant restaurantmodel.Restaurant
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create restaurant
	store := restaurantstorage.NewSQLStore(db)
	bizRestaurant := restaurantbusiness.NewCreateRestaurantBiz(store)

	if err := bizRestaurant.CreateRestaurant(&restaurantmodel.Restaurant{
		Name:        restaurant.Name,
		Email:       restaurant.Email,
		Password:    restaurant.Password,
		PhoneNumber: restaurant.PhoneNumber,
		PriceRange:  restaurant.PriceRange,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": restaurant})
}

// GetAllRestaurants all restaurant
func GetAllRestaurants(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	store := restaurantstorage.NewSQLStore(db)
	bizRestaurant := restaurantbusiness.NewListRestaurantBiz(store)

	restaurants, err := bizRestaurant.ListAllRestaurant()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}

// GetRestaurant a restaurant
func GetRestaurant(c *gin.Context) {
	id := c.Param("res-id")
	db := c.MustGet("db").(*gorm.DB)
	store := restaurantstorage.NewSQLStore(db)
	bizRestaurant := restaurantbusiness.NewGetRestaurantBiz(store)

	restaurant, err := bizRestaurant.GetRestaurant(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}

// UpdateRestaurant update
func UpdateRestaurant(c *gin.Context) {
	// TODO: Validate input
	var restaurant restaurantmodel.Restaurant

	id := c.Param("res-id")
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store := restaurantstorage.NewSQLStore(db)
	bizRestaurant := restaurantbusiness.NewUpdateRestaurantBiz(store)

	if err := bizRestaurant.UpdateRestaurant(
		id,
		restaurantmodel.Restaurant{
			Name:        restaurant.Name,
			PhoneNumber: restaurant.PhoneNumber,
			PriceRange:  restaurant.PriceRange,
		},
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "updated"})
}

// DeleteRestaurant a restaurant
func DeleteRestaurant(c *gin.Context) {
	id := c.Param("usr-id")
	db := c.MustGet("db").(*gorm.DB)
	store := restaurantstorage.NewSQLStore(db)
	bizRestaurant := restaurantbusiness.NewDeleteRestaurantBiz(store)

	if err := bizRestaurant.DeleteRestaurant(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "deleted"})
}
