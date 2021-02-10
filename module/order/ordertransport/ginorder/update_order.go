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

// UpdateOrder a order
func UpdateOrder(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var order ordermodel.OrderUpdate

		if err := c.ShouldBind(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		oid, err := common.FromBase58(c.Param("order-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := orderstorage.NewSQLStore(db)
		bizOrder := orderbusiness.NewUpdateOrderBiz(store)

		if err := bizOrder.UpdateOrder(c.Request.Context(), int(oid.GetLocalID()), &order); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
