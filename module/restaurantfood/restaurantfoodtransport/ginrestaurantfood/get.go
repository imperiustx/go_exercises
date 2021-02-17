package ginrestaurantfood

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodbusiness"
	"github.com/imperiustx/go_excercises/module/restaurantfood/restaurantfoodstorage"
)

// GetRestaurantFood a restaurantfood
func GetRestaurantFood(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := restaurantfoodstorage.NewSQLStore(db)
		bizRestaurantFood := restaurantfoodbusiness.NewGetRestaurantFoodBiz(store)

		rid, err := common.FromBase58(c.Param("rid"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fid, err := common.FromBase58(c.Param("fid"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		restaurantfood, err := bizRestaurantFood.GetRestaurantFood(
			c.Request.Context(),
			int(rid.GetLocalID()),
			int(fid.GetLocalID()),
		)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(restaurantfood))
	}
}
