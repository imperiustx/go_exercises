package ginimage

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/image/imagebusiness"
	"github.com/imperiustx/go_excercises/module/image/imagestorage"
)

// GetImage a image
func GetImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := imagestorage.NewSQLStore(db)

		bizImage := imagebusiness.NewGetImageBiz(store)
		image, err := bizImage.GetImage(c.Request.Context(), id)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(image))
	}
}
