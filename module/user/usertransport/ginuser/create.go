package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

// CreateUser a user
func CreateUser(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user usermodel.UserCreate
		if err := c.ShouldBindJSON(&user); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create user
		db := appCtx.GetDBConnection()
		store := userstorage.NewSQLStore(db)

		bizUser := userbusiness.NewCreateUserBiz(store)

		if err := bizUser.CreateNewUser(c.Request.Context(), &user); err != nil {
			panic(err)
		}
		user.GenUID(common.DBTypeUser, 1)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(user.FakeID))
	}
}
