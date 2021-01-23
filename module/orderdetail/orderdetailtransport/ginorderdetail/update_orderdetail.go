package ginorderdetail

import (
	"net/http"
	"strconv"

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
		idString := c.Param("ord-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := orderdetailstorage.NewSQLStore(db)
		bizOrderDetail := orderdetailbusiness.NewUpdateOrderDetailBiz(store)

		if err := bizOrderDetail.UpdateOrderDetail(c.Request.Context(), id, &orderdetail); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
