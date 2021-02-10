package ginorder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/orderbusiness"
	"github.com/imperiustx/go_excercises/module/order/orderstorage"
)

// GetOrder a order
func GetOrder(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("order-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := orderstorage.NewSQLStore(db)

		bizOrder := orderbusiness.NewGetOrderBiz(store)
		order, err := bizOrder.GetOrder(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		order.GenUID(common.DBTypeOrder, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(order))
	}
}
