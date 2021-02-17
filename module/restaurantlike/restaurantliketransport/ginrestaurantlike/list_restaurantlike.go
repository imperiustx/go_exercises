package ginrestaurantlike

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikebusiness"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikemodel"
	"github.com/imperiustx/go_excercises/module/restaurantlike/restaurantlikestorage"
)

// ListRestaurantLike a restaurantlike
func ListRestaurantLike(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			filter restaurantlikemodel.Filter
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
		if filter.RestaurantID != "" {
			rid, err := common.FromBase58(filter.RestaurantID)
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}

			filter.RestaurantID = strconv.Itoa(int(rid.GetLocalID()))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := restaurantlikestorage.NewSQLStore(db)
		bizRestaurantLike := restaurantlikebusiness.NewListRestaurantLikeBiz(store)

		restaurantlikes, err := bizRestaurantLike.ListRestaurantLike(c.Request.Context(), &filter, &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range restaurantlikes {
			restaurantlikes[i].GenUID(common.DBTypeUser, common.DBTypeRestaurant, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(restaurantlikes, paging, nil))
	}
}
