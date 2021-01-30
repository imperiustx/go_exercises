package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/appctx/hasher"
	"github.com/imperiustx/go_excercises/appctx/tokenprovider/jwt"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/userbusiness"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
	"github.com/imperiustx/go_excercises/module/user/userstorage"
)

func Login(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(err)
		}

		db := appCtx.GetDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbusiness.NewLoginBusiness(
			store,
			tokenProvider,
			md5,
			appctx.NewTokenConfig(),
		)
		account, err := business.Login(
			c.Request.Context(), 
			&loginUserData,
		)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
