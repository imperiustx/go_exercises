package ginrestaurantrating

import (
	"net/http"

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
		store := restaurantratingstorage.NewSQLStore(db)
		bizRestaurantRating := restaurantratingbusiness.NewUpdateRestaurantRatingBiz(store)

		rrid, err := common.FromBase58(c.Param("rr-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizRestaurantRating.UpdateRestaurantRating(
			c.Request.Context(),
			map[string]interface{}{"id": int(rrid.GetLocalID())},
			&restaurantrating); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
