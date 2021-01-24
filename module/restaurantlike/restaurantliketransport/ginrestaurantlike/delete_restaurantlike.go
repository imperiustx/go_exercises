package ginrestaurantlike

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikebusiness"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikestorage"
)

// DeleteRestaurantLike a restaurantlike
func DeleteRestaurantLike(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uidString := c.Param("uid")
		uid, err := strconv.Atoi(uidString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		ridString := c.Param("rid")
		rid, err := strconv.Atoi(ridString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := restaurantlikestorage.NewSQLStore(db)
		bizRestaurantLike := restaurantlikebusiness.NewDeleteRestaurantLikeBiz(store)

		if err := bizRestaurantLike.DeleteRestaurantLike(c.Request.Context(), uid, rid); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
