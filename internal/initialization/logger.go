package initialization

import (
	"go-ecommerce/global"
	"go-ecommerce/pkg/logger"

	"go.uber.org/zap"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)

	global.Logger.Info("Logger configure completed!", zap.String("ok", "success"))
}
