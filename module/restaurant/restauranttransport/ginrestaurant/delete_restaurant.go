package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantbusiness"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantstorage"
)

// DeleteRestaurant a restaurant
func DeleteRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		idString := c.Param("usr-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db := appCtx.GetDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		bizRestaurant := restaurantbusiness.NewDeleteRestaurantBiz(store)

		if err := bizRestaurant.DeleteRestaurant(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": "deleted"})
	}
}
