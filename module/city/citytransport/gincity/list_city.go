package gincity

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/city/citybusiness"
	"github.com/imperiustx/go_excercises/module/city/citystorage"
)

// ListCity a city
func ListCity(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := citystorage.NewSQLStore(db)

		bizCity := citybusiness.NewListCityBiz(store)
		cities, err := bizCity.ListAllCity(c.Request.Context(), &paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(cities, paging, nil))
	}
}
