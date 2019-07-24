package user

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"go_demo/handler"
	"go_demo/model"
	"go_demo/package/errno"
)

func Delete(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))

	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(ctx, nil, nil)
}
