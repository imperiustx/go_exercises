package ginrestaurantrating

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingbusiness"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingmodel"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingstorage"
)

// UpdateRestaurantRating a restaurantrating
func UpdateRestaurantRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		var restaurantrating restaurantratingmodel.RestaurantRatingUpdate

		if err := c.ShouldBindJSON(&restaurantrating); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		idString := c.Param("res-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantratingstorage.NewSQLStore(db)
		bizRestaurantRating := restaurantratingbusiness.NewUpdateRestaurantRatingBiz(store)

		if err := bizRestaurantRating.UpdateRestaurantRating(c.Request.Context(), id, &restaurantrating); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
