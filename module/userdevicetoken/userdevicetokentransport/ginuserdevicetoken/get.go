package ginuserdevicetoken

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenbusiness"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenstorage"
)

func Get(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := userdevicetokenstorage.NewSQLStore(db)
		bizUserDeviceToken := userdevicetokenbusiness.NewGetUserDeviceTokenBiz(store)

		uid, err := common.FromBase58(c.Param("udt-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userdevicetoken, err := bizUserDeviceToken.GetUserDeviceToken(
			c.Request.Context(),
			map[string]interface{}{"id": int(uid.GetLocalID())},
		)
		if err != nil {
			panic(err)
		}

		userdevicetoken.GenUID(common.DBTypeUserDeviceToken, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(userdevicetoken))
	}
}
