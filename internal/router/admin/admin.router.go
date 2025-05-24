package admin

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (r *AdminRouter) InitAdminRouter(rGroup *gin.RouterGroup) {
	// public
	adminRouterPublic := rGroup.Group("/admin")
	{
		adminRouterPublic.POST("/login")
		adminRouterPublic.POST("/otp")
	}

	// private
}