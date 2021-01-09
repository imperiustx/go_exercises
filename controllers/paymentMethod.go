package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/models"
	"gorm.io/gorm"
)

// CreatePaymentMethod new payment
func CreatePaymentMethod(c *gin.Context) {
	// TODO: Validate input
	var payment models.PaymentMethod
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create payment

	db.Create(&models.PaymentMethod{
		BankName: payment.BankName,
		UserName: payment.UserName,
		Number:   payment.Number,
	})

	c.JSON(http.StatusCreated, gin.H{"data": payment})
}

// GetAllPayments all payment
func GetAllPayments(c *gin.Context) {
	var payments []models.PaymentMethod

	db := c.MustGet("db").(*gorm.DB)

	db.Select("bank_name", "user_name", "number").Find(&payments)

	c.JSON(http.StatusOK, gin.H{"data": payments})
}

// GetPayment a payment
func GetPayment(c *gin.Context) {
	var payment models.PaymentMethod

	id := c.Param("pm-id")

	db := c.MustGet("db").(*gorm.DB)

	db.Where("id = ?", id).Select("bank_name", "user_name", "number").Find(&payment)

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

// UpdatePayment update
func UpdatePayment(c *gin.Context) {
	// TODO: Validate input
	var payment models.PaymentMethod

	id := c.Param("pm-id")

	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&payment).Where("id = ?", id).Updates(
		models.PaymentMethod{
			BankName: payment.BankName,
			UserName: payment.UserName,
			Number:   payment.Number,
		})

	c.JSON(http.StatusOK, gin.H{"data": "updated"})
}
