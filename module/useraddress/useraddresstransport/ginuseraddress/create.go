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

// CreateUserAddress a useraddress
func CreateUserAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var useraddress useraddressmodel.UserAddressCreate
		if err := c.ShouldBindJSON(&useraddress); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := useraddressstorage.NewSQLStore(db)
		bizUserAddress := useraddressbusiness.NewCreateUserAddressBiz(store)

		if err := bizUserAddress.CreateNewUserAddress(c.Request.Context(), &useraddress); err != nil {
			panic(err)
		}
		
		useraddress.GenUID(common.DBTypeUserAddress, 1)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(useraddress.FakeID))
	}
}
