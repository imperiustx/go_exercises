package ginimage

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/image/imagebusiness"
	"github.com/imperiustx/go_excercises/module/image/imagemodel"
	"github.com/imperiustx/go_excercises/module/image/imagestorage"
)

// CreateImage a image
func CreateImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var image imagemodel.ImageCreate
		if err := c.ShouldBindJSON(&image); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// Create image
		db := appCtx.GetDBConnection()
		store := imagestorage.NewSQLStore(db)

		bizImage := imagebusiness.NewCreateImageBiz(store)

		if err := bizImage.CreateNewImage(c.Request.Context(), &image); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(image))
	}
}
