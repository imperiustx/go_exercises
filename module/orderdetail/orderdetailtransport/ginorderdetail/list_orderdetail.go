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

// ListOrderDetail a orderdetail
func ListOrderDetail(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			filter orderdetailmodel.Filter
			order  common.OrderSort
		)
		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := orderdetailstorage.NewSQLStore(db)
		bizOrderDetail := orderdetailbusiness.NewListOrderDetailBiz(store)

		orderDetails, err := bizOrderDetail.ListAllOrderDetail(c.Request.Context(), &filter, &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range orderDetails {
			orderDetails[i].GenUID(common.DBTypeOrderDetail, common.DBTypeOrder, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(orderDetails, paging, nil))
	}
}
