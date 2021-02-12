package ginuseraddress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressbusiness"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressstorage"
)

// ListUserAddress a useraddress
func ListUserAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			filter useraddressmodel.Filter
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
		store := useraddressstorage.NewSQLStore(db)
		bizUserAddress := useraddressbusiness.NewListUserAddressBiz(store)

		useraddresses, err := bizUserAddress.ListAllUserAddress(c.Request.Context(), &filter, &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range useraddresses {
			useraddresses[i].GenUID(common.DBTypeUser, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(useraddresses, paging, nil))
	}
}
