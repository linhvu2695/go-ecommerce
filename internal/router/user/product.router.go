package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (r *ProductRouter) InitProductRouter(rGroup *gin.RouterGroup) {
	// public
	productRouterPublic := rGroup.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}
}
