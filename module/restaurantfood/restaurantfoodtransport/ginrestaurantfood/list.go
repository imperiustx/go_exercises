package ginrestaurantfood

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodbusiness"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodstorage"
)

// ListRestaurantFood a restaurantfood
func ListRestaurantFood(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			order  common.OrderSort
		)
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := restaurantfoodstorage.NewSQLStore(db)
		bizRestaurantFood := restaurantfoodbusiness.NewListRestaurantFoodBiz(store)

		restaurantfoods, err := bizRestaurantFood.ListAllRestaurantFood(
			c.Request.Context(),
			&paging,
			&order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range restaurantfoods {
			restaurantfoods[i].GenUID(common.DBTypeRestaurant, common.DBTypeFood, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(restaurantfoods, paging, nil))
	}
}
