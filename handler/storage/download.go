package storage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Download(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
