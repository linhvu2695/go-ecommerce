package middleware

import (
	"context"
	"go-ecommerce/global"
	"go-ecommerce/internal/constants"
	"go-ecommerce/pkg/utils/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.URL.Path
		global.Logger.Info("URI Request: ", zap.String("uri", uri))

		// Check headers authentication
		jwtToken, ok := auth.ExtractBearerToken(ctx)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "cannot extract bearer token",
			})
			return
		}

		// Validate JWT token by subject
		claims, err := auth.VerifyTokenSubject(jwtToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "invalid token",
			})
			return
		}

		// Update claims to context
		global.Logger.Info("claims:", zap.String("uuid", claims.Subject))
		c := context.WithValue(ctx.Request.Context(), constants.SUBJECT_UUID_KEY, claims.Subject)
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()

	}
}
