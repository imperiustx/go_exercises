package ginaddress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressbusiness"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
	"github.com/imperiustx/go_excercises/module/address/addressstorage"
)

// CreateAddress a address
func CreateAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var address addressmodel.AddressCreate
		if err := c.ShouldBindJSON(&address); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create address
		db := appCtx.GetDBConnection()
		store := addressstorage.NewSQLStore(db)

		bizAddress := addressbusiness.NewCreateAddressBiz(store)

		if err := bizAddress.CreateNewAddress(c.Request.Context(), &address); err != nil {
			panic(err)
		}
		address.GenUID(common.DBTypeAddress, 4)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(address.FakeID))
	}
}
