package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

func GetProfile(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("user-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := userstorage.NewSQLStore(db)
		bizUser := userbusiness.NewGetUserBiz(store)
		
		user, err := bizUser.GetUser(
			c.Request.Context(),
			map[string]interface{}{"id": int(uid.GetLocalID())},
		)
		if err != nil {
			panic(err)
		}

		user.Mask(true)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
