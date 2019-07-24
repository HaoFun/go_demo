package user

import (
	"github.com/gin-gonic/gin"

	"go_demo/handler"
	"go_demo/model"
	"go_demo/package/errno"
)

func Get(ctx *gin.Context) {
	username := ctx.Param("username")

	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(ctx, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(ctx, nil, user)
}