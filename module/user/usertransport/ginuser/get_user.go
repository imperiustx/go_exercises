package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
	"gorm.io/gorm"
)

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
