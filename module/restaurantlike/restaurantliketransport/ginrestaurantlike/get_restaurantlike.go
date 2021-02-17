package ginrestaurantlike

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikebusiness"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikestorage"
)

// GetRestaurantLike a restaurantlike
func GetRestaurantLike(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := restaurantlikestorage.NewSQLStore(db)
		bizRestaurantLike := restaurantlikebusiness.NewGetRestaurantLikeBiz(store)

		uid, err := common.FromBase58(c.Param("uid"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		rid, err := common.FromBase58(c.Param("rid"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		restaurantlike, err := bizRestaurantLike.GetRestaurantLike(
			c.Request.Context(),
			int(uid.GetLocalID()),
			int(rid.GetLocalID()),
		)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(restaurantlike))
	}
}
