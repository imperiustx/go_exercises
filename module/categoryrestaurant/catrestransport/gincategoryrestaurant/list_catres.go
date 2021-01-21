package gincategoryrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantbusiness"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantstorage"
)

// ListCategoryRestaurant a categoryrestaurant
func ListCategoryRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := categoryrestaurantstorage.NewSQLStore(db)

		bizCategoryRestaurant := categoryrestaurantbusiness.NewListCategoryRestaurantBiz(store)
		categoryrestaurants, err := bizCategoryRestaurant.ListAllCategoryRestaurant(c.Request.Context(), &paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(categoryrestaurants, paging, nil))
	}
}
