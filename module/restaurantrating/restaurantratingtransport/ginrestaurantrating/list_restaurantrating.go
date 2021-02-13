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

// ListRestaurantRating a restaurantrating
func ListRestaurantRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			filter restaurantratingmodel.Filter
			order  common.OrderSort
		)
		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := restaurantratingstorage.NewSQLStore(db)
		bizRestaurantRating := restaurantratingbusiness.NewListRestaurantRatingBiz(store)

		restaurantratings, err := bizRestaurantRating.ListAllRestaurantRating(c.Request.Context(), &filter, &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range restaurantratings {
			restaurantratings[i].GenUID(common.DBTypeRestaurantRating, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(restaurantratings, paging, nil))
	}
}
