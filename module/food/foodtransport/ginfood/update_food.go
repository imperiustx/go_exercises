package ginfood

import (
	"net/http"
	"strconv"

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
		idString := c.Param("food-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := foodstorage.NewSQLStore(db)
		bizFood := foodbusiness.NewUpdateFoodBiz(store)

		if err := bizFood.UpdateFood(c.Request.Context(), id, &food); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
