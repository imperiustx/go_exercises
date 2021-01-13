package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/module/paymentmethod/paymentmethodbusiness"
	"github.com/imperiustx/go_excercises/module/paymentmethod/paymentmethodmodel"
	"github.com/imperiustx/go_excercises/module/paymentmethod/paymentmethodstorage"
	"gorm.io/gorm"
)

// CreatePaymentMethod new payment
func CreatePaymentMethod(c *gin.Context) {
	// TODO: Validate input
	var payment paymentmethodmodel.PaymentMethod
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create payment
	store := paymentmethodstorage.NewSQLStore(db)
	bizPaymentMethod := paymentmethodbusiness.NewCreatePaymentMethodBiz(store)

	if err := bizPaymentMethod.CreatePaymentMethod(&paymentmethodmodel.PaymentMethod{
		BankName: payment.BankName,
		UserName: payment.UserName,
		Number:   payment.Number,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": payment})
}

// GetAllPayments all payment
func GetAllPayments(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	store := paymentmethodstorage.NewSQLStore(db)
	bizPaymentMethod := paymentmethodbusiness.NewListPaymentMethodBiz(store)

	methods, err := bizPaymentMethod.ListAllPaymentMethod()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": methods})
}

// GetPayment a payment
func GetPayment(c *gin.Context) {
	id := c.Param("pm-id")
	db := c.MustGet("db").(*gorm.DB)

	store := paymentmethodstorage.NewSQLStore(db)
	bizPaymentMethod := paymentmethodbusiness.NewGetPaymentMethodBiz(store)

	method, err := bizPaymentMethod.GetPaymentMethod(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": method})
}

// UpdatePayment update
func UpdatePayment(c *gin.Context) {
	// TODO: Validate input
	var payment paymentmethodmodel.PaymentMethod

	id := c.Param("pm-id")
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store := paymentmethodstorage.NewSQLStore(db)
	bizPaymentMethod := paymentmethodbusiness.NewUpdatePaymentMethodBiz(store)

	if err := bizPaymentMethod.UpdatePaymentMethod(
		id,
		paymentmethodmodel.PaymentMethod{
			BankName: payment.BankName,
			UserName: payment.UserName,
			Number:   payment.Number,
		},
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "updated"})
}
