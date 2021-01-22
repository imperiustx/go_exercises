package ginfoodlike

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikebusiness"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikestorage"
)

// DeleteFoodLike a foodlike
func DeleteFoodLike(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uidString := c.Param("uid")
		uid, err := strconv.Atoi(uidString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		fidString := c.Param("fid")
		fid, err := strconv.Atoi(fidString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := foodlikestorage.NewSQLStore(db)
		bizFoodLike := foodlikebusiness.NewDeleteFoodLikeBiz(store)

		if err := bizFoodLike.DeleteFoodLike(c.Request.Context(), uid, fid); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
