package admin

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(rGroup *gin.RouterGroup) {
	// private
	userRouterPrivate := rGroup.Group("/admin/user")
	{
		userRouterPrivate.POST("/validate")
	}
}