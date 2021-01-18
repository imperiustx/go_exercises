package gincity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citybusiness"
	"github.com/imperiustx/go_excercises/module/city/citymodel"
	"github.com/imperiustx/go_excercises/module/city/citystorage"
)

// CreateCity a city
func CreateCity(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var city citymodel.CityCreate
		if err := c.ShouldBindJSON(&city); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create city
		db := appCtx.GetDBConnection()
		store := citystorage.NewSQLStore(db)

		bizCity := citybusiness.NewCreateCityBiz(store)

		if err := bizCity.CreateNewCity(c.Request.Context(), &city); err != nil {
			panic(err)
		}
		city.GenUID(common.DBTypeCity, 3)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(city.FakeID))
	}
}
