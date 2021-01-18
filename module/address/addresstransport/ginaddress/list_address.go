package ginaddress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressbusiness"
	"github.com/imperiustx/go_excercises/module/address/addressstorage"
)

// ListAddress a address
func ListAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := addressstorage.NewSQLStore(db)

		bizAddress := addressbusiness.NewListAddressBiz(store)
		addresses, err := bizAddress.ListAllAddress(c.Request.Context(), &paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(addresses, paging, nil))
	}
}
