package ginorderdetail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailbusiness"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailstorage"
)

// GetOrderDetail a orderdetail
func GetOrderDetail(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := orderdetailstorage.NewSQLStore(db)
		bizOrderDetail := orderdetailbusiness.NewGetOrderDetailBiz(store)

		odid, err := common.FromBase58(c.Param("od-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		orderdetail, err := bizOrderDetail.GetOrderDetail(
			c.Request.Context(), 
			map[string]interface{}{"id": int(odid.GetLocalID())},
		)
		if err != nil {
			panic(err)
		}

		orderdetail.GenUID(common.DBTypeOrderDetail, common.DBTypeOrder, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(orderdetail))
	}
}
