package ginuser

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

// UpdateUser a user
func UpdateUser(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user usermodel.UserUpdate

		if err := c.ShouldBind(&user); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		idString := c.Param("usr-id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		bizUser := userbusiness.NewUpdateUserBiz(store)

		if err := bizUser.UpdateUser(c.Request.Context(), id, &user); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
