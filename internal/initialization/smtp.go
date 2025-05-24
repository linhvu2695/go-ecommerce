package initialization

import (
	"go-ecommerce/global"
	"net/smtp"

	"go.uber.org/zap"
)

func InitSmtp() {
	// SMTP configuration
	config := global.Config.Smtp

	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)

	global.SmtpAuth = &auth
	global.Logger.Info("SMTP configure completed!", zap.String("ok", "success"))
}
