package ginaddress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/address/addressbusiness"
	"github.com/imperiustx/go_excercises/module/address/addressstorage"
)

// GetAllAddresses a address
func GetAllAddresses(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()
		store := addressstorage.NewSQLStore(db)

		bizAddress := addressbusiness.NewListAddressBiz(store)
		addresses, err := bizAddress.ListAllAddress()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": addresses})
	}
}
