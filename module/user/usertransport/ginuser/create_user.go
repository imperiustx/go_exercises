package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

// CreateUser a user
func CreateUser(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user *usermodel.User
		db := appCtx.GetDBConnection()
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create user
		store := userstorage.NewSQLStore(db)
		bizUser := userbusiness.NewCreateUserBiz(store)

		if err := bizUser.CreateUser(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": &user})
	}
}
