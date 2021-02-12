package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantbusiness"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantstorage"
)

// ListRestaurant a restaurant
func ListRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			filter restaurantmodel.Filter
			order  common.OrderSort
		)

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := restaurantstorage.NewSQLStore(db)

		bizRestaurant := restaurantbusiness.NewListRestaurantBiz(store)
		restaurants, err := bizRestaurant.ListAllRestaurant(c.Request.Context(), &filter, &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range restaurants {
			restaurants[i].GenUID(common.DBTypeRestaurant, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(restaurants, paging, nil))
	}
}
