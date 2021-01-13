package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressbusiness"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressstorage"
	"gorm.io/gorm"
)

// CreateUserAddress new userAddress
func CreateUserAddress(c *gin.Context) {
	// TODO: Validate input
	var userAddress useraddressmodel.UserAddress
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&userAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create userAddress
	store := useraddressstorage.NewSQLStore(db)
	bizUserAddress := useraddressbusiness.NewCreateUserAddressBiz(store)

	if err := bizUserAddress.CreateUserAddress(&useraddressmodel.UserAddress{
		UserID:    userAddress.UserID,
		AddressID: userAddress.AddressID,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": userAddress})
}
