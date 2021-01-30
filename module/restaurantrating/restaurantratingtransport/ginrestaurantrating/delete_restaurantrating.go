package ginrestaurantrating

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingbusiness"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingstorage"
)

// DeleteRestaurantRating a restaurantrating
func DeleteRestaurantRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		idString := c.Param("res-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		db := appCtx.GetDBConnection()
		store := restaurantratingstorage.NewSQLStore(db)
		bizRestaurantRating := restaurantratingbusiness.NewDeleteRestaurantRatingBiz(store)

		if err := bizRestaurantRating.DeleteRestaurantRating(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
