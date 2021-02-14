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

// UpdateOrderDetail a orderdetail
func UpdateOrderDetail(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var orderdetail orderdetailmodel.OrderDetailUpdate

		if err := c.ShouldBind(&orderdetail); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := orderdetailstorage.NewSQLStore(db)
		bizOrderDetail := orderdetailbusiness.NewUpdateOrderDetailBiz(store)

		odid, err := common.FromBase58(c.Param("od-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizOrderDetail.UpdateOrderDetail(
			c.Request.Context(),
			map[string]interface{}{"id": int(odid.GetLocalID())},
			&orderdetail); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
