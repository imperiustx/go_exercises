package gincategoryrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantbusiness"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantstorage"
)

// UpdateCategoryRestaurant a categoryrestaurant
func UpdateCategoryRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var categoryrestaurant categoryrestaurantmodel.CategoryRestaurantUpdate

		if err := c.ShouldBind(&categoryrestaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
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

		store := categoryrestaurantstorage.NewSQLStore(db)
		bizCategoryRestaurant := categoryrestaurantbusiness.NewUpdateCategoryRestaurantBiz(store)

		if err := bizCategoryRestaurant.UpdateCategoryRestaurant(c.Request.Context(), cid, rid, &categoryrestaurant); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
