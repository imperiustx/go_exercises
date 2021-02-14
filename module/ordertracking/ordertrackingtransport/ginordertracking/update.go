package ginordertracking

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingbusiness"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingmodel"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingstorage"
)

// UpdateOrderTracking a ordertracking
func UpdateOrderTracking(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var ordertracking ordertrackingmodel.OrderTrackingUpdate

		if err := c.ShouldBind(&ordertracking); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := ordertrackingstorage.NewSQLStore(db)
		bizOrderTracking := ordertrackingbusiness.NewUpdateOrderTrackingBiz(store)

		oid, err := common.FromBase58(c.Param("ot-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := bizOrderTracking.UpdateOrderTracking(
			c.Request.Context(),
			map[string]interface{}{"id": int(oid.GetLocalID())},
			&ordertracking); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
