package gincategory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorybusiness"
	"github.com/imperiustx/go_excercises/module/category/categorystorage"
)

// GetCategory a category
func GetCategory(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		cid, err := common.FromBase58(c.Param("cat-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := categorystorage.NewSQLStore(db)

		bizCategory := categorybusiness.NewGetCategoryBiz(store)
		category, err := bizCategory.GetCategory(c.Request.Context(), int(cid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		category.GenUID(common.DBTypeCategory, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(category))
	}
}
