package initialization

import (
	"go-ecommerce/global"
	"go-ecommerce/internal/controller"
	"go-ecommerce/internal/router"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares

	// routers
	adminRouter := router.RouterGroupApp.Admin
	userRouter := router.RouterGroupApp.User

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
	}
	{
		adminRouter.InitAdminRouter(v1)
		adminRouter.InitUserRouter(v1)
	}
	{
		userRouter.InitProductRouter(v1)
		userRouter.InitUserRouter(v1)
	}

	return r
}
