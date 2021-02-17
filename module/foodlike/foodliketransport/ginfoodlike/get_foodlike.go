package ginfoodlike

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikebusiness"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikestorage"
)

// GetFoodLike a foodlike
func GetFoodLike(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := foodlikestorage.NewSQLStore(db)
		bizFoodLike := foodlikebusiness.NewGetFoodLikeBiz(store)

		uid, err := common.FromBase58(c.Param("uid"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fid, err := common.FromBase58(c.Param("fid"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		foodlike, err := bizFoodLike.GetFoodLike(c.Request.Context(),
			int(uid.GetLocalID()),
			int(fid.GetLocalID()),
		)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(foodlike))
	}
}
