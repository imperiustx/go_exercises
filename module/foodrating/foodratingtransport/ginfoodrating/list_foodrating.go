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

// ListFoodRating a foodrating
func ListFoodRating(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			filter foodratingmodel.Filter
			order  common.OrderSort
		)
		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := foodratingstorage.NewSQLStore(db)
		bizFoodRating := foodratingbusiness.NewListFoodRatingBiz(store)

		foodRatings, err := bizFoodRating.ListAllFoodRating(c.Request.Context(), &filter, &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range foodRatings {
			foodRatings[i].GenUID(common.DBTypeFoodRating, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(foodRatings, paging, nil))
	}
}
