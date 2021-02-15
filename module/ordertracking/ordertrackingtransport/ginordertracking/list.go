package ginordertracking

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingbusiness"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingstorage"
)

// ListOrderTracking a ordertracking
func ListOrderTracking(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			paging common.Paging
			order  common.OrderSort
		)
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&order); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := ordertrackingstorage.NewSQLStore(db)
		bizOrderTracking := ordertrackingbusiness.NewListOrderTrackingBiz(store)

		orderTrackings, err := bizOrderTracking.ListAllOrderTracking(c.Request.Context(), &paging, &order)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range orderTrackings {
			orderTrackings[i].GenUID(common.DBTypeOrderTracking, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(orderTrackings, paging, nil))
	}
}
