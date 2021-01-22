package gincategoryrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantbusiness"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantstorage"
)

// DeleteCategoryRestaurant a categoryrestaurant
func DeleteCategoryRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		cidString := c.Param("cid")
		cid, err := strconv.Atoi(cidString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		ridString := c.Param("rid")
		rid, err := strconv.Atoi(ridString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := categoryrestaurantstorage.NewSQLStore(db)
		bizCategoryRestaurant := categoryrestaurantbusiness.NewDeleteCategoryRestaurantBiz(store)

		if err := bizCategoryRestaurant.DeleteCategoryRestaurant(c.Request.Context(), cid, rid); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
