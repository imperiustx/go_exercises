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

// UpdateFoodRating a foodrating
func UpdateFoodRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var foodrating foodratingmodel.FoodRatingUpdate

		if err := c.ShouldBind(&foodrating); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := foodratingstorage.NewSQLStore(db)
		bizFoodRating := foodratingbusiness.NewUpdateFoodRatingBiz(store)

		frid, err := common.FromBase58(c.Param("fr-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizFoodRating.UpdateFoodRating(
			c.Request.Context(),
			map[string]interface{}{"id": int(frid.GetLocalID())},
			&foodrating); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
