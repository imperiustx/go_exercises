package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
	"gorm.io/gorm"
)

// CreateUser new user
func CreateUser(c *gin.Context) {
	// TODO: Validate input
	var user usermodel.User
	db := c.MustGet("db").(*gorm.DB)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	store := userstorage.NewSQLStore(db)
	bizUser := userbusiness.NewCreateUserBiz(store)

	if err := bizUser.CreateUser(&usermodel.User{
		FullName:    user.FullName,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// GetAllUsers all users
func GetAllUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	store := userstorage.NewSQLStore(db)
	bizUser := userbusiness.NewListUserBiz(store)
	users, err := bizUser.ListAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUser a user
func GetUser(c *gin.Context) {
	id := c.Param("usr-id")
	db := c.MustGet("db").(*gorm.DB)
	store := userstorage.NewSQLStore(db)
	bizUser := userbusiness.NewGetUserBiz(store)

	user, err := bizUser.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// UpdateUser update
func UpdateUser(c *gin.Context) {
	// TODO: Validate input
	var user usermodel.User

	id := c.Param("usr-id")
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store := userstorage.NewSQLStore(db)
	bizUser := userbusiness.NewUpdateUserBiz(store)

	if err := bizUser.UpdateUser(
		id,
		usermodel.User{
			FullName:    user.FullName,
			PhoneNumber: user.PhoneNumber,
		},
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "updated"})
}

// DeleteUser a user
func DeleteUser(c *gin.Context) {
	id := c.Param("usr-id")
	db := c.MustGet("db").(*gorm.DB)
	store := userstorage.NewSQLStore(db)
	bizUser := userbusiness.NewDeleteUserBiz(store)

	if err := bizUser.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "deleted"})
}
