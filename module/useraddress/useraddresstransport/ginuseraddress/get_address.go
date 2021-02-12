package ginuseraddress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressbusiness"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressstorage"
)

// GetUserAddress a useraddress
func GetUserAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := useraddressstorage.NewSQLStore(db)
		bizUserAddress := useraddressbusiness.NewGetUserAddressBiz(store)

		uaid, err := common.FromBase58(c.Param("ua-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		useraddress, err := bizUserAddress.GetUserAddress(
			c.Request.Context(),
			map[string]interface{}{"id": int(uaid.GetLocalID())})
		if err != nil {
			panic(err)
		}

		useraddress.GenUID(common.DBTypeUserAddress, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(useraddress))
	}
}
