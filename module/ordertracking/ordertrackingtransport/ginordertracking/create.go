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

// CreateOrderTracking a ordertracking
func CreateOrderTracking(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var ordertracking ordertrackingmodel.OrderTrackingCreate
		if err := c.ShouldBindJSON(&ordertracking); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create ordertracking
		db := appCtx.GetDBConnection()
		store := ordertrackingstorage.NewSQLStore(db)

		bizOrderTracking := ordertrackingbusiness.NewCreateOrderTrackingBiz(store)

		if err := bizOrderTracking.CreateNewOrderTracking(c.Request.Context(), &ordertracking); err != nil {
			panic(err)
		}
		ordertracking.GenUID(common.DBTypeOrderTracking, 1)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(ordertracking.FakeID))
	}
}
