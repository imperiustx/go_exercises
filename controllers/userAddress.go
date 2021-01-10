package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/models"
	"gorm.io/gorm"
)

// CreateUserAddress new userAddress
func CreateUserAddress(c *gin.Context) {
	// TODO: Validate input
	var userAddress models.UserAddress
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&userAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create userAddress

	db.Create(&models.UserAddress{
		UserID: userAddress.UserID,
		AddressID: userAddress.AddressID,
	})

	c.JSON(http.StatusCreated, gin.H{"data": userAddress})
}
