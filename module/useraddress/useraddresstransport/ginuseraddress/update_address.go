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

// UpdateUserAddress a useraddress
func UpdateUserAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var useraddress useraddressmodel.UserAddressUpdate

		if err := c.ShouldBind(&useraddress); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := useraddressstorage.NewSQLStore(db)
		bizUserAddress := useraddressbusiness.NewUpdateUserAddressBiz(store)

		uaid, err := common.FromBase58(c.Param("ua-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizUserAddress.UpdateUserAddress(
			c.Request.Context(),
			map[string]interface{}{"id": int(uaid.GetLocalID())},
			&useraddress); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
