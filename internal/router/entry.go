package router

import (
	"go-ecommerce/internal/router/admin"
	"go-ecommerce/internal/router/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Admin admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
