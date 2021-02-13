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

		db := appCtx.GetDBConnection()
		store := categorystorage.NewSQLStore(db)
		bizCategory := categorybusiness.NewGetCategoryBiz(store)

		cid, err := common.FromBase58(c.Param("user-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		category, err := bizCategory.GetCategory(
			c.Request.Context(),
			map[string]interface{}{"id": int(cid.GetLocalID())},
		)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(category))
	}
}
