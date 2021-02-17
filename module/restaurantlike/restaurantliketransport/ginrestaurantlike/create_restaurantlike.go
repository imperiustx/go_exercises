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

// CreateRestaurantLike a restaurantlike
func CreateRestaurantLike(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		var data restaurantlikemodel.RestaurantLikeCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := restaurantlikestorage.NewSQLStore(db)
		bizRestaurantLike := restaurantlikebusiness.NewCreateRestaurantLikeBiz(store)

		uid, err := common.FromBase58(data.UserID)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		rid, err := common.FromBase58(data.RestaurantID)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.RestaurantID = strconv.Itoa(int(rid.GetLocalID()))
		data.UserID = strconv.Itoa(int(uid.GetLocalID()))

		if err := bizRestaurantLike.CreateNewRestaurantLike(
			c.Request.Context(),
			&data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated,
			common.SimpleSuccessResponse(
				map[string]string{"data": "success"},
			),
		)
	}
}
