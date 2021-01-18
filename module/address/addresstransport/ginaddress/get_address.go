package ginaddress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressbusiness"
	"github.com/imperiustx/go_excercises/module/address/addressstorage"
)

// GetAddress a address
func GetAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("add-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := addressstorage.NewSQLStore(db)

		bizAddress := addressbusiness.NewGetAddressBiz(store)
		address, err := bizAddress.GetAddress(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		address.GenUID(common.DBTypeAddress, 4)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(address))
	}
}
