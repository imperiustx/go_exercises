package ginaddress

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/address/addressbusiness"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
	"github.com/imperiustx/go_excercises/module/address/addressstorage"
)

// UpdateAddress a address
func UpdateAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO: Validate input
		var address addressmodel.Address

		id := c.Param("usr-id")
		db := appCtx.GetDBConnection()

		if err := c.ShouldBindJSON(&address); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := addressstorage.NewSQLStore(db)
		bizAddress := addressbusiness.NewUpdateAddressBiz(store)

		if err := bizAddress.UpdateAddress(
			id,
			addressmodel.Address{
				FullAddress: address.FullAddress,
				Latitude:    address.Latitude,
				Longitude:   address.Longitude,
			},
		); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("address %s updated", id)})
	}
}
