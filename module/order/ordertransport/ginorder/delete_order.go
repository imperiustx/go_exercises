package ginorder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/order/orderbusiness"
	"github.com/imperiustx/go_excercises/module/order/orderstorage"
)

// DeleteOrder a order
func DeleteOrder(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		oid, err := common.FromBase58(c.Param("order-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := orderstorage.NewSQLStore(db)
		bizOrder := orderbusiness.NewDeleteOrderBiz(store)

		if err := bizOrder.DeleteOrder(c.Request.Context(), int(oid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
