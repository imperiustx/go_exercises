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

// DeleteImage a image
func DeleteImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		idString := c.Param("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := imagestorage.NewSQLStore(db)
		bizImage := imagebusiness.NewDeleteImageBiz(store)

		if err := bizImage.DeleteImage(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(map[string]int{"data": 1}))
	}
}
