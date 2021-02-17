package gincart

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartbusiness"
	"github.com/imperiustx/go_excercises/module/cart/cartmodel"
	"github.com/imperiustx/go_excercises/module/cart/cartstorage"
)

func Create(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		var data cartmodel.CartCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetDBConnection()
		store := cartstorage.NewSQLStore(db)
		repo := cartbusiness.NewCreateBusiness(store)

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

		if err := repo.Create(c.Request.Context(), &data); err != nil {
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
