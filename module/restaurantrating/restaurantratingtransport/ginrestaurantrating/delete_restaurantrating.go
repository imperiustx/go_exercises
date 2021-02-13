package ginrestaurantrating

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingbusiness"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingstorage"
)

// DeleteRestaurantRating a restaurantrating
func DeleteRestaurantRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := restaurantratingstorage.NewSQLStore(db)
		bizRestaurantRating := restaurantratingbusiness.NewDeleteRestaurantRatingBiz(store)

		rrid, err := common.FromBase58(c.Param("rr-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizRestaurantRating.DeleteRestaurantRating(
			c.Request.Context(),
			map[string]interface{}{"id": int(rrid.GetLocalID())},
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
