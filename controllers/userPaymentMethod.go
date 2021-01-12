package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/models"
	"gorm.io/gorm"
)

// CreateUserPaymentMethod new userPaymentMethod
func CreateUserPaymentMethod(c *gin.Context) {
	// TODO: Validate input
	var userPaymentMethod models.UserPaymentMethod
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&userPaymentMethod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create userPaymentMethod

	db.Create(&models.UserPaymentMethod{
		UserID:          userPaymentMethod.UserID,
		PaymentMethodID: userPaymentMethod.PaymentMethodID,
	})

	c.JSON(http.StatusCreated, gin.H{"data": userPaymentMethod})
}
