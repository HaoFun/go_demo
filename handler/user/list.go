package user

import (
	"github.com/gin-gonic/gin"

	"go_demo/handler"
	"go_demo/package/errno"
	userService "go_demo/service/user"
)

func List(ctx *gin.Context) {
	var request ListRequest

	if err := ctx.Bind(&request); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}

	infos, count, err := userService.ListUser(request.Username, request.Offset, request.Limit)
	if err != nil {
		handler.SendResponse(ctx, err, nil)
		return
	}

	handler.SendResponse(ctx, nil, ListResponse{
		TotalCount: count,
		UserList: infos,
	})
}