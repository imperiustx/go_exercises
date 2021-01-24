package ginrestaurantlike

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikebusiness"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikestorage"
)

// ListRestaurantLike a restaurantlike
func ListRestaurantLike(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := restaurantlikestorage.NewSQLStore(db)

		bizRestaurantLike := restaurantlikebusiness.NewListRestaurantLikeBiz(store)
		restaurantlikes, err := bizRestaurantLike.ListAllRestaurantLike(c.Request.Context(), &paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(restaurantlikes, paging, nil))
	}
}
