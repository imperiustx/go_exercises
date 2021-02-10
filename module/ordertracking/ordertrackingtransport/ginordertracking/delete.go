package ginordertracking

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingbusiness"
	"github.com/imperiustx/go_excercises/module/ordertracking/ordertrackingstorage"
)

// DeleteOrderTracking a ordertracking
func DeleteOrderTracking(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		oid, err := common.FromBase58(c.Param("ot-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := ordertrackingstorage.NewSQLStore(db)
		bizOrderTracking := ordertrackingbusiness.NewDeleteOrderTrackingBiz(store)

		if err := bizOrderTracking.DeleteOrderTracking(c.Request.Context(), int(oid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
