package ginaddress

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/address/addressbusiness"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
	"github.com/imperiustx/go_excercises/module/address/addressstorage"
)

// CreateAddress a address
func CreateAddress(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var address *addressmodel.Address
		db := appCtx.GetDBConnection()
		if err := c.ShouldBindJSON(&address); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create address
		store := addressstorage.NewSQLStore(db)
		bizAddress := addressbusiness.NewCreateAddressBiz(store)

		if err := bizAddress.CreateAddress(&address); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": &address})
	}
}
