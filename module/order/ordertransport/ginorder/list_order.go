package ginorder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/orderbusiness"
	"github.com/imperiustx/go_excercises/module/order/orderstorage"
)

// ListOrder a order
func ListOrder(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			order  common.OrderSort
		)
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := orderstorage.NewSQLStore(db)
		bizOrder := orderbusiness.NewListOrderBiz(store)

		orders, err := bizOrder.ListAllOrder(c.Request.Context(), &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range orders {
			orders[i].GenUID(common.DBTypeOrder, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(orders, paging, nil))
	}
}
