package gincity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citybusiness"
	"github.com/imperiustx/go_excercises/module/city/citystorage"
)

// GetCity a city
func GetCity(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("city-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := citystorage.NewSQLStore(db)

		bizCity := citybusiness.NewGetCityBiz(store)
		city, err := bizCity.GetCity(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		city.GenUID(common.DBTypeCity, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(city))
	}
}
