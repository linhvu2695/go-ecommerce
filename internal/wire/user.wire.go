//go:build wireinject

package wire

import (
	c "go-ecommerce/internal/controller/user"
	s "go-ecommerce/internal/service/user"

	"github.com/google/wire"
)

func InitLoginRouterHandler() (*c.LoginController, error) {
	wire.Build(
		s.UserLoginService,
		c.NewLoginController,
	)

	return new(c.LoginController), nil
}

func InitRegisterRouterHandler() (*c.RegisterController, error) {
	wire.Build(
		s.UserRegisterService,
		c.NewRegisterController,
	)

	return new(c.RegisterController), nil
}
