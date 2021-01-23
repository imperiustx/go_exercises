package ginrestaurantrating

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingbusiness"
	"github.com/imperiustx/go_excercises/module/restaurantrating/restaurantratingstorage"
)

// ListRestaurantRating a restaurantrating
func ListRestaurantRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := restaurantratingstorage.NewSQLStore(db)

		bizRestaurantRating := restaurantratingbusiness.NewListRestaurantRatingBiz(store)
		restaurantratings, err := bizRestaurantRating.ListAllRestaurantRating(c.Request.Context(), &paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(restaurantratings, paging, nil))
	}
}
