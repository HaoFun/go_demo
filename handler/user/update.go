package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	"go_demo/handler"
	"go_demo/model"
	"go_demo/package/errno"
	"go_demo/package/util"
)

func Update(ctx *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetRequestID(ctx)})

	userId, _ := strconv.Atoi(ctx.Param("id"))

	var user model.UserModel

	if err := ctx.Bind(&user); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}

	user.Id = uint64(userId)

	if err := user.Encrypt(); err != nil {
		handler.SendResponse(ctx, errno.ErrEncrypt, nil)
		return
	}

	if err := user.Update(); err != nil {
		handler.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(ctx, nil, nil)
}