package ginorderdetail

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailbusiness"
	"github.com/imperiustx/go_excercises/module/orderdetail/orderdetailstorage"
)

// DeleteOrderDetail a orderdetail
func DeleteOrderDetail(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetDBConnection()
		store := orderdetailstorage.NewSQLStore(db)
		bizOrderDetail := orderdetailbusiness.NewDeleteOrderDetailBiz(store)

		odid, err := common.FromBase58(c.Param("od-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := bizOrderDetail.DeleteOrderDetail(
			c.Request.Context(),
			map[string]interface{}{"id": int(odid.GetLocalID())},
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
