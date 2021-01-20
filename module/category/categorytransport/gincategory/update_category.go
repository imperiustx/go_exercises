package gincategory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorybusiness"
	"github.com/imperiustx/go_excercises/module/category/categorymodel"
	"github.com/imperiustx/go_excercises/module/category/categorystorage"
)

// UpdateCategory a category
func UpdateCategory(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var category categorymodel.CategoryUpdate

		if err := c.ShouldBind(&category); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		idString := c.Param("cat-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSQLStore(db)
		bizCategory := categorybusiness.NewUpdateCategoryBiz(store)

		if err := bizCategory.UpdateCategory(c.Request.Context(), id, &category); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
