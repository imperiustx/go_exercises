package gincart

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartbusiness"
	"github.com/imperiustx/go_excercises/module/cart/cartstorage"
)

func Get(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uidString := c.Param("u-id")
		uid, err := strconv.Atoi(uidString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fidString := c.Param("f-id")
		fid, err := strconv.Atoi(fidString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := cartstorage.NewSQLStore(db)
		bizCart := cartbusiness.NewGetCartBiz(store)

		cart, err := bizCart.GetCart(
			c.Request.Context(),
			uid,
			fid,
		)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(cart))
	}
}
