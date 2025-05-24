package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractBearerToken(ctx *gin.Context) (string, bool) {
	authHeader := ctx.GetHeader("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer "), true
	}
	return "", false
}
