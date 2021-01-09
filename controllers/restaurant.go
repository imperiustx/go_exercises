package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/models"
	"gorm.io/gorm"
)

// CreateRestaurant new restaurant
func CreateRestaurant(c *gin.Context) {
	// TODO: Validate input
	var restaurant models.Restaurant
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create restaurant

	db.Create(&models.Restaurant{
		Name:        restaurant.Name,
		Email:       restaurant.Email,
		Password:    restaurant.Password,
		PhoneNumber: restaurant.PhoneNumber,
		PriceRange:  restaurant.PriceRange,
	})

	c.JSON(http.StatusCreated, gin.H{"data": restaurant})
}

// GetAllRestaurants all restaurant
func GetAllRestaurants(c *gin.Context) {
	var restaurants []models.Restaurant

	db := c.MustGet("db").(*gorm.DB)

	db.Select("name", "email", "phone_number").Find(&restaurants)

	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}

// GetRestaurant a restaurant
func GetRestaurant(c *gin.Context) {
	var restaurant models.Restaurant

	id := c.Param("res-id")

	db := c.MustGet("db").(*gorm.DB)

	db.Where("id = ?", id).Select("name", "phone_number").Find(&restaurant)

	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}

// UpdateRestaurant update
func UpdateRestaurant(c *gin.Context) {
	// TODO: Validate input
	var restaurant models.Restaurant

	id := c.Param("res-id")

	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&restaurant).Where("id = ?", id).Updates(
		models.Restaurant{
			Name:        restaurant.Name,
			PhoneNumber: restaurant.PhoneNumber,
			PriceRange:  restaurant.PriceRange,
		})

	c.JSON(http.StatusOK, gin.H{"data": "updated"})
}
