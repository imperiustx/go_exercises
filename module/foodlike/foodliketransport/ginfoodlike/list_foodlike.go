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

// ListFoodLike a foodlike
func ListFoodLike(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			filter foodlikemodel.Filter
			order  common.OrderSort
		)
		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if filter.UserID != "" {
			uid, err := common.FromBase58(filter.UserID)
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}

			filter.UserID = strconv.Itoa(int(uid.GetLocalID()))
		}
		if filter.FoodID != "" {
			fid, err := common.FromBase58(filter.FoodID)
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}

			filter.FoodID = strconv.Itoa(int(fid.GetLocalID()))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := foodlikestorage.NewSQLStore(db)
		bizFoodLike := foodlikebusiness.NewListFoodLikeBiz(store)

		foodlikes, err := bizFoodLike.ListAllFoodLike(c.Request.Context(), &filter, &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range foodlikes {
			foodlikes[i].GenUID(common.DBTypeUser, common.DBTypeFood, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(foodlikes, paging, nil))
	}
}
