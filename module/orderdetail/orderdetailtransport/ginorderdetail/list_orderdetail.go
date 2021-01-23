package ginorderdetail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailbusiness"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailstorage"
)

// ListOrderDetail a orderdetail
func ListOrderDetail(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := orderdetailstorage.NewSQLStore(db)

		bizOrderDetail := orderdetailbusiness.NewListOrderDetailBiz(store)
		orderdetails, err := bizOrderDetail.ListAllOrderDetail(c.Request.Context(), &paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(orderdetails, paging, nil))
	}
}
