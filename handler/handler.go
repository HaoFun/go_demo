package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go_demo/package/errno"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func SendResponse(ctx *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Message: message,
		Data: data,
	})
}