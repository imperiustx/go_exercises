package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/module/userpaymentmethod/userpaymentmethodbusiness"
	"github.com/imperiustx/go_excercises/module/userpaymentmethod/userpaymentmethodmodel"
	"github.com/imperiustx/go_excercises/module/userpaymentmethod/userpaymentmethodstorage"
	"gorm.io/gorm"
)

// CreateUserPaymentMethod new userPaymentMethod
func CreateUserPaymentMethod(c *gin.Context) {
	// TODO: Validate input
	var userPaymentMethod userpaymentmethodmodel.UserPaymentMethod
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&userPaymentMethod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create userPaymentMethod
	store := userpaymentmethodstorage.NewSQLStore(db)
	bizUserPaymentMethod := userpaymentmethodbusiness.NewCreateUserPaymentMethodBiz(store)

	if err := bizUserPaymentMethod.CreateUserPaymentMethod(&userpaymentmethodmodel.UserPaymentMethod{
		UserID:          userPaymentMethod.UserID,
		PaymentMethodID: userPaymentMethod.PaymentMethodID,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": userPaymentMethod})
}
