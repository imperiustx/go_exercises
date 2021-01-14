package ginuser

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

// UpdateUser a user
func UpdateUser(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO: Validate input
		var user usermodel.User

		id := c.Param("usr-id")
		db := appCtx.GetDBConnection()

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

		c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("user %s updated", id)})
	}
}
