package gincategory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorybusiness"
	"github.com/imperiustx/go_excercises/module/category/categorymodel"
	"github.com/imperiustx/go_excercises/module/category/categorystorage"
)

// CreateCategory a category
func CreateCategory(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var category categorymodel.CategoryCreate
		if err := c.ShouldBindJSON(&category); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create category
		db := appCtx.GetDBConnection()
		store := categorystorage.NewSQLStore(db)

		bizCategory := categorybusiness.NewCreateCategoryBiz(store)

		if err := bizCategory.CreateNewCategory(c.Request.Context(), &category); err != nil {
			panic(err)
		}
		category.GenUID(common.DBTypeCategory, 1)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(category.FakeID))
	}
}
