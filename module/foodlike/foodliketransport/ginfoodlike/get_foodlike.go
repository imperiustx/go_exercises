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

// GetFoodLike a foodlike
func GetFoodLike(appCtx appctx.AppContext) func(c *gin.Context) {
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

		bizFoodLike := foodlikebusiness.NewGetFoodLikeBiz(store)
		foodlike, err := bizFoodLike.GetFoodLike(c.Request.Context(), uid, fid)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(foodlike))
	}
}
