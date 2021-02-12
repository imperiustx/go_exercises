package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantbusiness"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantstorage"
)

// DeleteRestaurant a restaurant
func DeleteRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		rid, err := common.FromBase58(c.Param("res-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		bizRestaurant := restaurantbusiness.NewDeleteRestaurantBiz(store)

		if err := bizRestaurant.DeleteRestaurant(
			c.Request.Context(),
			map[string]interface{}{"id": int(rid.GetLocalID())}); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
