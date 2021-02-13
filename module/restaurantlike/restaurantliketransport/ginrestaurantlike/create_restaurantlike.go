package ginrestaurantlike

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikebusiness"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikestorage"
)

// CreateRestaurantLike a restaurantlike
func CreateRestaurantLike(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var restaurantlike restaurantlikemodel.RestaurantLikeCreate
		if err := c.ShouldBindJSON(&restaurantlike); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create restaurantlike
		db := appCtx.GetDBConnection()
		store := restaurantlikestorage.NewSQLStore(db)

		bizRestaurantLike := restaurantlikebusiness.NewCreateRestaurantLikeBiz(store)

		if err := bizRestaurantLike.CreateNewRestaurantLike(
			c.Request.Context(),
			&restaurantlike); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(restaurantlike))
	}
}
