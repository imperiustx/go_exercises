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
		rid, err := common.FromBase58(c.Param("res-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := restaurantratingstorage.NewSQLStore(db)

		bizRestaurantRating := restaurantratingbusiness.NewGetRestaurantRatingBiz(store)
		restaurantrating, err := bizRestaurantRating.GetRestaurantRating(c.Request.Context(), int(rid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		restaurantrating.GenUID(common.DBTypeRestaurantRating, 2)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(restaurantrating))
	}
}
