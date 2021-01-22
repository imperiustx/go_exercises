package ginfoodlike

import (
	"net/http"

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
		var foodlike foodlikemodel.FoodLikeCreate
		if err := c.ShouldBindJSON(&foodlike); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create foodlike
		db := appCtx.GetDBConnection()
		store := foodlikestorage.NewSQLStore(db)

		bizFoodLike := foodlikebusiness.NewCreateFoodLikeBiz(store)

		if err := bizFoodLike.CreateNewFoodLike(c.Request.Context(), &foodlike); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(foodlike))
	}
}
