package gincart

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/cart/cartbusiness"
	"github.com/imperiustx/go_excercises/module/cart/cartstorage"
)

func List(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := cartstorage.NewSQLStore(db)

		bizCart := cartbusiness.NewListCartBiz(store)

		carts, err := bizCart.ListAllCart(c.Request.Context(), &paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(carts, paging, nil))
	}
}
