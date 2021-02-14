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

func List(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			filter userdevicetokenmodel.Filter
			order  common.OrderSort
		)
		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := userdevicetokenstorage.NewSQLStore(db)
		bizUserDeviceToken := userdevicetokenbusiness.NewListUserDeviceTokenBiz(store)

		userdevicetokens, err := bizUserDeviceToken.ListAllUserDeviceToken(c.Request.Context(), &filter, &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range userdevicetokens {
			userdevicetokens[i].GenUID(common.DBTypeUserDeviceToken, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(userdevicetokens, paging, nil))
	}
}
