package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantbusiness"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantstorage"
)

// GetRestaurant a restaurant
func GetRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("res-id")
		db := appCtx.GetDBConnection()
		store := restaurantstorage.NewSQLStore(db)

		bizRestaurant := restaurantbusiness.NewGetRestaurantBiz(store)
		restaurant, err := bizRestaurant.GetRestaurant(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": restaurant})
	}
}
