package ginrestaurantrating

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingbusiness"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingstorage"
)

// GetRestaurantRating a restaurantrating
func GetRestaurantRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := restaurantratingstorage.NewSQLStore(db)
		bizRestaurantRating := restaurantratingbusiness.NewGetRestaurantRatingBiz(store)

		rrid, err := common.FromBase58(c.Param("rr-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		
		restaurantrating, err := bizRestaurantRating.GetRestaurantRating(
			c.Request.Context(),
			map[string]interface{}{"id": int(rrid.GetLocalID())})
		if err != nil {
			panic(err)
		}

		restaurantrating.GenUID(common.DBTypeRestaurantRating, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(restaurantrating))
	}
}
