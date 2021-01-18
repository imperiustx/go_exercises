package gincity

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citybusiness"
	"github.com/imperiustx/go_excercises/module/city/citymodel"
	"github.com/imperiustx/go_excercises/module/city/citystorage"
)

// UpdateCity a city
func UpdateCity(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var city citymodel.CityUpdate

		if err := c.ShouldBind(&city); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		idString := c.Param("city-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := citystorage.NewSQLStore(db)
		bizCity := citybusiness.NewUpdateCityBiz(store)

		if err := bizCity.UpdateCity(c.Request.Context(), id, &city); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
