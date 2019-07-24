package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func GetShortId() (string, error) {
	return shortid.Generate()
}

func GetRequestID(ctx *gin.Context) string {
	value, ok := ctx.Get("X-Request-Id")

	if !ok {
		return ""
	}

	if requestId, ok := value.(string); ok {
		return requestId
	}

	return ""
}