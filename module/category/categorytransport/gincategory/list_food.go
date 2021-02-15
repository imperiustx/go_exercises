package gincategory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorybusiness"
	"github.com/imperiustx/go_excercises/module/category/categorystorage"
)

func ListFood(appCtx appctx.AppContext) func(c *gin.Context) {
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
		store := categorystorage.NewSQLStore(db)
		bizCategory := categorybusiness.NewListFoodByCategoryIDBiz(store)

		cidString := c.Param("cat-id")
		cid, err := strconv.Atoi(cidString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		foods, err := bizCategory.ListAllFoodByCategoryID(c.Request.Context(), cid, &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range foods {
			foods[i].GenUID(common.DBTypeFood, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(foods, paging, nil))
	}
}
