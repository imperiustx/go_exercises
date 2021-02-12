package ginrestaurant

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantbusiness"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantstorage"
)

// CreateRestaurant a restaurant
func CreateRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetDBConnection()
		var restaurant restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&restaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		bizRestaurant := restaurantbusiness.NewCreateRestaurantBiz(store)
		fmt.Println("<><><><>", restaurant)
		if err := bizRestaurant.CreateNewRestaurant(c.Request.Context(), &restaurant); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(restaurant.FakeID))
	}
}
