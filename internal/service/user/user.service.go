package user

import (
	"context"
	"go-ecommerce/global"
	"go-ecommerce/internal/database"
	"go-ecommerce/internal/service/user/impl"
	"go-ecommerce/internal/service/user/models"
)

type (
	IUserLogin interface {
		Login(ctx context.Context, in *models.LoginInput) (out models.LoginOutput, err error)
		Logout(ctx context.Context, token string) error
		Setup2FA(ctx context.Context, in *models.Setup2FAInput) (int, error)
		Verify2FA(ctx context.Context, in *models.Verify2FAInput) (int, error)
	}

	IUserRegister interface {
		Register(ctx context.Context, in *models.RegisterInput) (out models.RegisterOutput, err error)
		VerifyOTP(ctx context.Context, in *models.VerifyOTPInput) (out models.VerifyOTPOutput, err error)
	}

	IUserInfo interface {
		GetUserInfo(ctx context.Context, id int) (int, error)
		UpdateUserInfo(ctx context.Context, userInfo map[string]interface{}) error
	}

	IUserAdmin interface {
		GetUserByID(ctx context.Context, id int) (int, error)
		DeleteUser(ctx context.Context, id int) error
	}
)

var (
	userLoginService    IUserLogin
	userRegisterService IUserRegister
	userInfoService     IUserInfo
	userAdminService    IUserAdmin
)

func UserLoginService() IUserLogin {
	if userLoginService == nil {
		r := database.New(global.Db)
		userLoginService = impl.NewUserLogin(r)
	}
	return userLoginService
}

func UserRegisterService() IUserRegister {
	if userRegisterService == nil {
		r := database.New(global.Db)
		userRegisterService = impl.NewUserRegister(r)
	}
	return userRegisterService
}

func UserInfoService() IUserInfo {
	if userInfoService == nil {
		panic("User info service not initialized")
	}
	return userInfoService
}

func UserAdminService() IUserAdmin {
	if userAdminService == nil {
		panic("User admin service not initialized")
	}
	return userAdminService
}
