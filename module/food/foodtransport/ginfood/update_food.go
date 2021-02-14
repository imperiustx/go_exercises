package ginfood

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodbusiness"
	"github.com/imperiustx/go_excercises/module/food/foodmodel"
	"github.com/imperiustx/go_excercises/module/food/foodstorage"
)

// UpdateFood a food
func UpdateFood(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var food foodmodel.FoodUpdate

		if err := c.ShouldBind(&food); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := foodstorage.NewSQLStore(db)
		bizFood := foodbusiness.NewUpdateFoodBiz(store)

		fid, err := common.FromBase58(c.Param("food-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizFood.UpdateFood(
			c.Request.Context(),
			map[string]interface{}{"id": int(fid.GetLocalID())},
			&food); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
