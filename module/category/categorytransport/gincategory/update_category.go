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

// UpdateCategory a category
func UpdateCategory(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var category categorymodel.CategoryUpdate

		if err := c.ShouldBind(&category); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := categorystorage.NewSQLStore(db)
		bizCategory := categorybusiness.NewUpdateCategoryBiz(store)

		cid, err := common.FromBase58(c.Param("cat-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizCategory.UpdateCategory(
			c.Request.Context(),
			map[string]interface{}{"id": int(cid.GetLocalID())},
			&category); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
