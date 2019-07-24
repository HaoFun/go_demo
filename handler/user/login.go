package user

import (
	"github.com/gin-gonic/gin"

	"go_demo/handler"
	"go_demo/model"
	"go_demo/package/auth"
	"go_demo/package/errno"
	"go_demo/package/token"
)

func Login(ctx *gin.Context) {
    var request model.UserModel
    if err := ctx.Bind(&request); err != nil {
    	handler.SendResponse(ctx, errno.ErrBind, nil)
    	return
	}

    d, err := model.GetUser(request.Username)
    if err != nil {
    	handler.SendResponse(ctx, errno.ErrUserNotFound, nil)
    	return
	}

    if err := auth.Compare(d.Password, request.Password); err != nil {
    	handler.SendResponse(ctx, errno.ErrPasswordIncorrect, nil)
    	return
	}

    t, err := token.Sign(
    	ctx,
    	token.Context{
    		ID: d.Id,
    		Username: d.Username,
    	},
    	"",
    )

    if err != nil {
    	handler.SendResponse(ctx, errno.ErrToken, nil)
    	return
	}

    handler.SendResponse(
    	ctx,
    	nil,
    	model.Token{
    		Token:"Bearer " + t,
    	},
    )
}