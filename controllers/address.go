package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/module/address/addressbusiness"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
	"github.com/imperiustx/go_excercises/module/address/addressstorage"
	"gorm.io/gorm"
)

// CreateAddress new address
func CreateAddress(c *gin.Context) {
	// TODO: Validate input
	var address addressmodel.Address
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create address
	store := addressstorage.NewSQLStore(db)
	bizAddress := addressbusiness.NewCreateAddressBiz(store)

	if err := bizAddress.CreateAddress(&addressmodel.Address{
		FullAddress: address.FullAddress,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": address})
}

// GetAllAddresses all address
func GetAllAddresses(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	store := addressstorage.NewSQLStore(db)
	bizAddress := addressbusiness.NewListAddressBiz(store)

	addresses, err := bizAddress.ListAllAddress()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": addresses})
}

// GetAddress a address
func GetAddress(c *gin.Context) {
	id := c.Param("add-id")
	db := c.MustGet("db").(*gorm.DB)
	store := addressstorage.NewSQLStore(db)
	bizAddress := addressbusiness.NewGetAddressBiz(store)

	address, err := bizAddress.GetAddress(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": address})
}

// UpdateAddress update
func UpdateAddress(c *gin.Context) {
	// TODO: Validate input
	var address addressmodel.Address

	id := c.Param("add-id")
	db := c.MustGet("db").(*gorm.DB)
	
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store := addressstorage.NewSQLStore(db)
	bizAddress := addressbusiness.NewUpdateAddressBiz(store)
	if err := bizAddress.UpdateAddress(
		id, 
		addressmodel.Address{
			FullAddress: address.FullAddress,
			Latitude:    address.Latitude,
			Longitude:   address.Longitude,
		},
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "updated"})
}
