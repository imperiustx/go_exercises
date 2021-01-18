package ginaddress

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressbusiness"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
	"github.com/imperiustx/go_excercises/module/address/addressstorage"
)

// UpdateAddress a address
func UpdateAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var address addressmodel.AddressUpdate

		if err := c.ShouldBind(&address); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		idString := c.Param("usr-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := addressstorage.NewSQLStore(db)
		bizAddress := addressbusiness.NewUpdateAddressBiz(store)

		if err := bizAddress.UpdateAddress(c.Request.Context(), id, &address); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
