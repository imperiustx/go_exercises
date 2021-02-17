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

// CreateFood a food
func CreateFood(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var food foodmodel.FoodCreate
		if err := c.ShouldBindJSON(&food); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := foodstorage.NewSQLStore(db)
		bizFood := foodbusiness.NewCreateFoodBiz(store)

		rid, err := common.FromBase58(food.RestaurantID)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		cid, err := common.FromBase58(food.CategoryID)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		food.RestaurantID = strconv.Itoa(int(rid.GetLocalID()))
		food.CategoryID = strconv.Itoa(int(cid.GetLocalID()))

		if err := bizFood.CreateNewFood(c.Request.Context(), &food); err != nil {
			panic(err)
		}

		food.GenUID(common.DBTypeFood, 1)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(food.FakeID))
	}
}
