package ginuserdevicetoken

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenbusiness"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenmodel"
	"github.com/imperiustx/go_excercises/module/userdevicetoken/userdevicetokenstorage"
)

func Create(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data userdevicetokenmodel.UserDeviceTokenCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetDBConnection()
		store := userdevicetokenstorage.NewSQLStore(db)
		repo := userdevicetokenbusiness.NewCreateUserDeviceTokenBusiness(store)

		if err := repo.CreateUserDeviceToken(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DBTypeUserDeviceToken, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.ID))
	}
}
