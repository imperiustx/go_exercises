package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

func Delete(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := userstorage.NewSQLStore(db)
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		bizUser := userbusiness.NewDeleteUserBiz(store, requester)

		uid, err := common.FromBase58(c.Param("user-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizUser.DeleteUser(
			c.Request.Context(),
			map[string]interface{}{"id": int(uid.GetLocalID())},
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
