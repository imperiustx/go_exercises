package ginuseraddress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressbusiness"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressstorage"
)

// DeleteUserAddress a useraddress
func DeleteUserAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := useraddressstorage.NewSQLStore(db)
		bizUserAddress := useraddressbusiness.NewDeleteUserAddressBiz(store)

		uaid, err := common.FromBase58(c.Param("ua-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizUserAddress.DeleteUserAddress(
			c.Request.Context(),
			map[string]interface{}{"id": int(uaid.GetLocalID())},
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
