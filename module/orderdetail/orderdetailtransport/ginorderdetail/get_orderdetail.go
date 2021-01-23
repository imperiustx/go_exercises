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
		uid, err := common.FromBase58(c.Param("ord-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := orderdetailstorage.NewSQLStore(db)

		bizOrderDetail := orderdetailbusiness.NewGetOrderDetailBiz(store)
		orderdetail, err := bizOrderDetail.GetOrderDetail(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		orderdetail.GenUID(common.DBTypeOrderDetail, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(orderdetail))
	}
}
