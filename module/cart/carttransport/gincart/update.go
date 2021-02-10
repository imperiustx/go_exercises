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

func Update(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var cart cartmodel.CartUpdate

		if err := c.ShouldBind(&cart); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
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

		store := cartstorage.NewSQLStore(db)
		bizCart := cartbusiness.NewUpdateCartBiz(store)

		if err := bizCart.UpdateCart(c.Request.Context(), uid, fid, &cart); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
