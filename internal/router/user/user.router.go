package user

import (
	"go-ecommerce/global"
	"go-ecommerce/internal/middleware"
	"go-ecommerce/internal/wire"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(rGroup *gin.RouterGroup) {
	registerController, err := wire.InitRegisterRouterHandler()
	if err != nil {
		global.Logger.Error("failed to initialize register controller", zap.Error(err))
		return
	}

	loginController, err := wire.InitLoginRouterHandler()
	if err != nil {
		global.Logger.Error("failed to initialize login controller", zap.Error(err))
		return
	}

	// public
	userRouterPublic := rGroup.Group("/user")
	{
		userRouterPublic.POST("/register", registerController.Register)
		userRouterPublic.POST("/verify_otp", registerController.VerifyOTP)
		userRouterPublic.POST("/login", loginController.Login)
		userRouterPublic.POST("/otp")
	}

	// private
	userRouterPrivate := rGroup.Group("/user")
	userRouterPrivate.Use(middleware.AuthMiddleware())
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/2fa/setup", loginController.Setup2FA)
		userRouterPrivate.POST("/2fa/verify", loginController.Verify2FA)
	}
}
