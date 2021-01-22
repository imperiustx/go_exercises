package ginfoodrating

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingbusiness"
	"github.com/imperiustx/go_excercises/module/foodrating/foodratingstorage"
)

// DeleteFoodRating a foodrating
func DeleteFoodRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		idString := c.Param("fr-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := foodratingstorage.NewSQLStore(db)
		bizFoodRating := foodratingbusiness.NewDeleteFoodRatingBiz(store)

		if err := bizFoodRating.DeleteFoodRating(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
