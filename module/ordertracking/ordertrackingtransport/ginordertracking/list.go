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
			paging        common.Paging
			ordertracking common.OrderSort
		)
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		if err := c.ShouldBind(&ordertracking); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := ordertrackingstorage.NewSQLStore(db)
		// requester := c.MustGet(common.CurrentUser).(common.Requester)
		bizOrderTracking := ordertrackingbusiness.NewListOrderTrackingBiz(store)
		ordertrackings, err := bizOrderTracking.ListAllOrderTracking(c.Request.Context(), &paging, &ordertracking)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(ordertrackings, paging, nil))
	}
}
