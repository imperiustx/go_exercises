package ginordertracking

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingbusiness"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingstorage"
)

// GetOrderTracking a ordertracking
func GetOrderTracking(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("ot-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := ordertrackingstorage.NewSQLStore(db)

		bizOrderTracking := ordertrackingbusiness.NewGetOrderTrackingBiz(store)
		ordertracking, err := bizOrderTracking.GetOrderTracking(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		ordertracking.GenUID(common.DBTypeOrderTracking, 1)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(ordertracking))
	}
}
