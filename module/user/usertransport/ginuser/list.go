package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

// ListUser a user
func ListUser(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBindJSON(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		store := userstorage.NewSQLStore(db)

		bizUser := userbusiness.NewListUserBiz(store)
		users, err := bizUser.ListAllUser(c.Request.Context(), &paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for i := range users {
			users[i].GenUID(common.DBTypeUser, 1)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, nil))
	}
}
