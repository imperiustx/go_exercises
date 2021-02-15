package ginfoodrating

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingbusiness"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingmodel"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingstorage"
)

// CreateFoodRating a foodrating
func CreateFoodRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var foodrating foodratingmodel.FoodRatingCreate
		if err := c.ShouldBindJSON(&foodrating); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := foodratingstorage.NewSQLStore(db)
		bizFoodRating := foodratingbusiness.NewCreateFoodRatingBiz(store)

		if err := bizFoodRating.CreateNewFoodRating(c.Request.Context(), &foodrating); err != nil {
			panic(err)
		}
		
		foodrating.GenUID(common.DBTypeFoodRating, 1)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(foodrating.FakeID))
	}
}
