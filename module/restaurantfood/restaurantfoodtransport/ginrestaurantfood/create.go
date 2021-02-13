package ginrestaurantfood

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodbusiness"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodmodel"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodstorage"
)

// CreateRestaurantFood a restaurantfood
func CreateRestaurantFood(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var restaurantfood restaurantfoodmodel.RestaurantFoodCreate
		if err := c.ShouldBindJSON(&restaurantfood); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := restaurantfoodstorage.NewSQLStore(db)
		bizRestaurantFood := restaurantfoodbusiness.NewCreateRestaurantFoodBiz(store)

		if err := bizRestaurantFood.CreateNewRestaurantFood(
			c.Request.Context(),
			&restaurantfood); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(restaurantfood))
	}
}
