package gincategory

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/category/categorybusiness"
	"github.com/imperiustx/go_excercises/module/category/categorystorage"
)

// DeleteCategory a category
func DeleteCategory(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := categorystorage.NewSQLStore(db)
		bizCategory := categorybusiness.NewDeleteCategoryBiz(store)

		cid, err := common.FromBase58(c.Param("cat-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizCategory.DeleteCategory(
			c.Request.Context(),
			map[string]interface{}{"id": int(cid.GetLocalID())},
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
