package userbusiness

import (
	"context"

	"github.com/imperiustx/go_excercises/appctx"
	"github.com/imperiustx/go_excercises/appctx/tokenprovider"
	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type TokenConfig interface {
	GetAtExp() int
	GetRtExp() int
}

type loginBusiness struct {
	appCtx        appctx.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	tkCfg         TokenConfig
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, tkCfg TokenConfig) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		tkCfg:         tkCfg,
	}
}

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := business.storeUser.FindUser(
		ctx,
		map[string]interface{}{"email": data.Email},
	)
	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	passHashed := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserID: user.ID,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(
		payload,
		business.tkCfg.GetAtExp(),
	)
	if err != nil {
		return nil, common.ErrCannotGenerateToken("access", err)
	}

	refreshToken, err := business.tokenProvider.Generate(
		payload,
		business.tkCfg.GetRtExp(),
	)
	if err != nil {
		return nil, common.ErrCannotGenerateToken("refresh", err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
