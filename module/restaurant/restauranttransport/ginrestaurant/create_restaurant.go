package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantbusiness"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantstorage"
)

// CreateRestaurant a restaurant
func CreateRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var restaurant *restaurantmodel.Restaurant
		db := appCtx.GetDBConnection()
		if err := c.ShouldBindJSON(&restaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Create restaurant
		store := restaurantstorage.NewSQLStore(db)
		bizRestaurant := restaurantbusiness.NewCreateRestaurantBiz(store)

		if err := bizRestaurant.CreateRestaurant(&restaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": &restaurant})
	}
}
