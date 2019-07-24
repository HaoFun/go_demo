package middleware

import (
	"github.com/gin-gonic/gin"

	"go_demo/handler"
	"go_demo/package/errno"
	"go_demo/package/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, err := token.ParseRequest(ctx); err != nil {
			handler.SendResponse(ctx, errno.ErrTokenInvalid, nil)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}