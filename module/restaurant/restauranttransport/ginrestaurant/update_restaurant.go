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

// UpdateRestaurant a restaurant
func UpdateRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		var restaurant restaurantmodel.RestaurantUpdate

		if err := c.ShouldBindJSON(&restaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		rid, err := common.FromBase58(c.Param("res-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		bizRestaurant := restaurantbusiness.NewUpdateRestaurantBiz(store)

		if err := bizRestaurant.UpdateRestaurant(
			c.Request.Context(),
			map[string]interface{}{"id": int(rid.GetLocalID())},
			&restaurant); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
