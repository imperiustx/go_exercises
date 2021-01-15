package ginrestaurant

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantbusiness"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantstorage"
)

// UpdateRestaurant a restaurant
func UpdateRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO: Validate input
		var restaurant restaurantmodel.Restaurant

		id := c.Param("usr-id")
		db := appCtx.GetDBConnection()

		if err := c.ShouldBindJSON(&restaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		bizRestaurant := restaurantbusiness.NewUpdateRestaurantBiz(store)

		if err := bizRestaurant.UpdateRestaurant(
			id,
			restaurantmodel.Restaurant{
				Name:        restaurant.Name,
				PhoneNumber: restaurant.PhoneNumber,
			},
		); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("restaurant %s updated", id)})
	}
}
