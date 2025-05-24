package main

import (
	"fmt"
	"go-ecommerce/internal/initialization"

	_ "go-ecommerce/cmd/swag/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation Go-Ecommerce
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Linh Vu
// @contact.url    https://github.com/linhvu2695
// @contact.email  linhvu2695@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @schema           http
func main() {
	fmt.Println("Starting...")

	r := initialization.Run()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8082")
}
