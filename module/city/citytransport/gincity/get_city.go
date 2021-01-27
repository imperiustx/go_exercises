package gincity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citybusiness"
	"github.com/imperiustx/go_excercises/module/city/citystorage"
)

// GetCity a city
func GetCity(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		cid, err := strconv.Atoi(c.Param("city-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := appCtx.GetDBConnection()
		store := citystorage.NewSQLStore(db)

		bizCity := citybusiness.NewGetCityBiz(store)
		city, err := bizCity.GetCity(c.Request.Context(), cid)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(city))
	}
}
