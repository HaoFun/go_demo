package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"

	"go_demo/model"
	"go_demo/package/util"
	"go_demo/handler"
	"go_demo/package/errno"
)

func Create(ctx *gin.Context) {
	log.Info(
		"User Create function called.",
		lager.Data{"X-Request-Id": util.GetRequestID(ctx)},
	)

	var request CreateRequest

	if err := ctx.Bind(&request); err != nil {
		handler.SendResponse(ctx, errno.ErrBind, nil)
		return
	}

	user := model.UserModel{
		Username: request.Username,
		Password: request.Password,
	}

	if err := user.Validate(); err != nil {
		handler.SendResponse(ctx, errno.ErrValidation, nil)
		return
	}

	if err := user.Encrypt(); err != nil {
		handler.SendResponse(ctx, errno.ErrEncrypt, nil)
		return
	}

	if err := user.Create(); err != nil {
		handler.SendResponse(ctx, errno.ErrDatabase, nil)
		return
	}

	response := CreateResponse{
		Username: user.Username,
	}

	handler.SendResponse(ctx, nil, response)
}