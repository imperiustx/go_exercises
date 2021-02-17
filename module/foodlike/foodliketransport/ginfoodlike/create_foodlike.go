package ginfoodlike

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikebusiness"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikemodel"
	"github.com/imperiustx/go_excercises/module/foodlike/foodlikestorage"
)

// CreateFoodLike a foodlike
func CreateFoodLike(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		var data foodlikemodel.FoodLikeCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := foodlikestorage.NewSQLStore(db)
		bizFoodLike := foodlikebusiness.NewCreateFoodLikeBiz(store)

		uid, err := common.FromBase58(data.UserID)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		fid, err := common.FromBase58(data.FoodID)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.FoodID = strconv.Itoa(int(fid.GetLocalID()))
		data.UserID = strconv.Itoa(int(uid.GetLocalID()))

		if err := bizFoodLike.CreateNewFoodLike(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(
			http.StatusCreated,
			common.SimpleSuccessResponse(
				map[string]string{"data": "success"},
			),
		)
	}
}
