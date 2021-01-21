package gincategoryrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantbusiness"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantmodel"
	"github.com/imperiustx/go_excercises/module/categoryrestaurant/categoryrestaurantstorage"
)

// CreateCategoryRestaurant a categoryrestaurant
func CreateCategoryRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var categoryrestaurant categoryrestaurantmodel.CategoryRestaurantCreate
		if err := c.ShouldBindJSON(&categoryrestaurant); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create categoryrestaurant
		db := appCtx.GetDBConnection()
		store := categoryrestaurantstorage.NewSQLStore(db)

		bizCategoryRestaurant := categoryrestaurantbusiness.NewCreateCategoryRestaurantBiz(store)

		if err := bizCategoryRestaurant.CreateNewCategoryRestaurant(c.Request.Context(), &categoryrestaurant); err != nil {
			panic(err)
		}
		categoryrestaurant.GenUID(common.DBTypeCategoryRestaurant, 1)

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(categoryrestaurant.FakeID))
	}
}
