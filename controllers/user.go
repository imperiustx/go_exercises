package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/models"
	"gorm.io/gorm"
)

// CreateUser new user
func CreateUser(c *gin.Context) {
	// TODO: Validate input
	var user models.User
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user

	db.Create(&models.User{
		FullName:    user.FullName,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
	})

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// GetAllUsers all user
func GetAllUsers(c *gin.Context) {
	var users []models.User

	db := c.MustGet("db").(*gorm.DB)

	db.Select("full_name", "email", "phone_number").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUser a user
func GetUser(c *gin.Context) {
	var user models.User

	id := c.Param("usr-id")

	db := c.MustGet("db").(*gorm.DB)

	/*
	
	TODO: convert below
		SELECT full_name, email, phone_number, full_address, bank_name, number
		FROM users u
		INNER JOIN user_addresses ON u.id = user_addresses.id
		INNER JOIN addresses a ON user_addresses.address_id = a.id
		INNER JOIN user_payment_methods ON u.id = user_payment_methods.user_id
		INNER JOIN payment_methods p ON user_payment_methods.id = p.id
		WHERE u.id = 1;
	*/

	db.Where("id = ?", id).Select("full_name", "email", "phone_number").Find(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser update
func UpdateUser(c *gin.Context) {
	// TODO: Validate input
	var user models.User

	id := c.Param("usr-id")

	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&user).Where("id = ?", id).Updates(
		models.User{
			FullName:    user.FullName,
			PhoneNumber: user.PhoneNumber,
		})

		c.JSON(http.StatusOK, gin.H{"data": "updated"})
}
