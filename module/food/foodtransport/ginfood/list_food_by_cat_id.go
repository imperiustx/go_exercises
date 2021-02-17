package ginfood

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/food/foodbusiness"
	"github.com/imperiustx/go_excercises/module/food/foodstorage"
)

func ListFoodByCategoryID(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			order  common.OrderSort
		)
		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := foodstorage.NewSQLStore(db)
		bizCategory := foodbusiness.NewListFoodByCategoryIDBiz(store)

		cid, err := common.FromBase58(c.Param("cat-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		foods, err := bizCategory.ListAllFoodByCategoryID(
			c.Request.Context(),
			int(cid.GetLocalID()),
			&paging,
			&order,
		)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range foods {
			foods[i].GenUID(common.DBTypeFood, common.DBTypeRestaurant, common.DBTypeCategory, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(foods, paging, nil))
	}
}
