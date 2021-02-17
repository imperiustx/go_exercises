package gincart

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartbusiness"
	"github.com/imperiustx/go_excercises/module/cart/cartstorage"
)

func Get(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := cartstorage.NewSQLStore(db)
		bizCart := cartbusiness.NewGetCartBiz(store)

		uid, err := common.FromBase58(c.Param("u-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fid, err := common.FromBase58(c.Param("f-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		cart, err := bizCart.GetCart(
			c.Request.Context(),
			int(uid.GetLocalID()),
			int(fid.GetLocalID()),
		)
		if err != nil {
			panic(err)
		}

		cart.GenUID(common.DBTypeUser, common.DBTypeFood, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(cart))
	}
}
