package ginorder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/orderbusiness"
	"github.com/imperiustx/go_excercises/module/order/ordermodel"
	"github.com/imperiustx/go_excercises/module/order/orderstorage"
)

// CreateOrder a order
func CreateOrder(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var order ordermodel.OrderCreate
		if err := c.ShouldBindJSON(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create order
		db := appCtx.GetDBConnection()
		store := orderstorage.NewSQLStore(db)

		bizOrder := orderbusiness.NewCreateOrderBiz(store)

		if err := bizOrder.CreateNewOrder(c.Request.Context(), &order); err != nil {
			panic(err)
		}
		order.GenUID(common.DBTypeOrder, 1)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(order.FakeID))
	}
}
