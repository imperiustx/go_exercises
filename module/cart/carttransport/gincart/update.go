package gincart

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartbusiness"
	"github.com/imperiustx/go_excercises/module/cart/cartmodel"
	"github.com/imperiustx/go_excercises/module/cart/cartstorage"
)

func Update(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var cart cartmodel.CartUpdate

		if err := c.ShouldBind(&cart); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := cartstorage.NewSQLStore(db)
		bizCart := cartbusiness.NewUpdateCartBiz(store)

		uid, err := common.FromBase58(c.Param("u-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fid, err := common.FromBase58(c.Param("f-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizCart.UpdateCart(
			c.Request.Context(),
			int(uid.GetLocalID()),
			int(fid.GetLocalID()),
			&cart,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
