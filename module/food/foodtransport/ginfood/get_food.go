package ginfood

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodbusiness"
	"github.com/imperiustx/go_excercises/module/food/foodstorage"
)

// GetFood a food
func GetFood(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := foodstorage.NewSQLStore(db)
		bizFood := foodbusiness.NewGetFoodBiz(store)

		fid, err := common.FromBase58(c.Param("food-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		food, err := bizFood.GetFood(
			c.Request.Context(),
			map[string]interface{}{"id": int(fid.GetLocalID())},
		)
		if err != nil {
			panic(err)
		}

		food.GenUID(common.DBTypeFood, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(food))
	}
}
