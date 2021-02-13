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

// CreateRestaurantRating a restaurantrating
func CreateRestaurantRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()
		var restaurantrating restaurantratingmodel.RestaurantRatingCreate

		if err := c.ShouldBindJSON(&restaurantrating); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantratingstorage.NewSQLStore(db)
		bizRestaurantRating := restaurantratingbusiness.NewCreateRestaurantRatingBiz(store)

		if err := bizRestaurantRating.CreateNewRestaurantRating(
			c.Request.Context(),
			&restaurantrating); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(restaurantrating.FakeID))
	}
}
