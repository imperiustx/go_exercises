package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

// GetUser a user
func GetUser(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("usr-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := userstorage.NewSQLStore(db)

		bizUser := userbusiness.NewGetUserBiz(store)
		user, err := bizUser.GetUser(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		user.GenUID(common.DBTypeUser, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
