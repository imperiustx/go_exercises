package gincategory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorybusiness"
	"github.com/imperiustx/go_excercises/module/category/categorystorage"
)

// ListCategory a category
func ListCategory(appCtx appctx.AppContext) func(c *gin.Context) {
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
		bizCategory := categorybusiness.NewListCategoryBiz(store)

		categories, err := bizCategory.ListAllCategory(c.Request.Context(), &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range categories {
			categories[i].GenUID(common.DBTypeCategory, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(categories, paging, nil))
	}
}
