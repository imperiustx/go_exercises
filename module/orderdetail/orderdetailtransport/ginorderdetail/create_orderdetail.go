package ginorderdetail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailbusiness"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailmodel"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailstorage"
)

// CreateOrderDetail a orderdetail
func CreateOrderDetail(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var orderdetail orderdetailmodel.OrderDetailCreate
		if err := c.ShouldBindJSON(&orderdetail); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create orderdetail
		db := appCtx.GetDBConnection()
		store := orderdetailstorage.NewSQLStore(db)

		bizOrderDetail := orderdetailbusiness.NewCreateOrderDetailBiz(store)

		if err := bizOrderDetail.CreateNewOrderDetail(c.Request.Context(), &orderdetail); err != nil {
			panic(err)
		}
		orderdetail.GenUID(common.DBTypeOrderDetail, 1)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(orderdetail.FakeID))
	}
}
