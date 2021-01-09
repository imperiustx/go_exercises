package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/models"
	"gorm.io/gorm"
)

// CreateAddress new address
func CreateAddress(c *gin.Context) {
	// TODO: Validate input
	var address models.Address
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create address

	db.Create(&models.Address{
		FullAddress: address.FullAddress,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
	})

	c.JSON(http.StatusCreated, gin.H{"data": address})
}

// GetAllAddresses all address
func GetAllAddresses(c *gin.Context) {
	var addresses []models.Address

	db := c.MustGet("db").(*gorm.DB)

	db.Select("full_address").Find(&addresses)

	c.JSON(http.StatusOK, gin.H{"data": addresses})
}

// GetAddress a address
func GetAddress(c *gin.Context) {
	var address models.Address

	id := c.Param("add-id")

	db := c.MustGet("db").(*gorm.DB)

	db.Where("id = ?", id).Select("full_address").Find(&address)

	c.JSON(http.StatusOK, gin.H{"data": address})
}

// UpdateAddress update
func UpdateAddress(c *gin.Context) {
	// TODO: Validate input
	var address models.Address

	id := c.Param("add-id")

	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&address).Where("id = ?", id).Updates(
		models.Address{
			FullAddress: address.FullAddress,
			Latitude:    address.Latitude,
			Longitude:   address.Longitude,
		})

	c.JSON(http.StatusOK, gin.H{"data": "updated"})
}
