package ginfoodrating

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingbusiness"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingstorage"
)

// DeleteFoodRating a foodrating
func DeleteFoodRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := foodratingstorage.NewSQLStore(db)
		bizFoodRating := foodratingbusiness.NewDeleteFoodRatingBiz(store)

		frid, err := common.FromBase58(c.Param("fr-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizFoodRating.DeleteFoodRating(
			c.Request.Context(),
			map[string]interface{}{"id": int(frid.GetLocalID())},
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
