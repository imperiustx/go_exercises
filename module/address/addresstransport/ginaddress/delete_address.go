package ginaddress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/address/addressbusiness"
	"github.com/imperiustx/go_excercises/module/address/addressstorage"
)

// DeleteAddress a address
func DeleteAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("usr-id")
		db := appCtx.GetDBConnection()
		store := addressstorage.NewSQLStore(db)
		bizAddress := addressbusiness.NewDeleteAddressBiz(store)

		if err := bizAddress.DeleteAddress(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "deleted"})
	}
}
