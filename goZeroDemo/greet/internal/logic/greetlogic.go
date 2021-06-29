package logic

import (
	"context"

	"pam/greet/internal/svc"
	"pam/greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GreetLogic {
	return GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Message: "hello go-zero",
	}, nil
}

func (l *GreetLogic) Bye(req types.Request) (*types.ByeResponse, error) {
	Message := "ok"
	Code := "200"
	Data := map[string]string{"name": "Japan"}

	return &types.ByeResponse{
		Message: Message,
		Code:    Code,
		Data:    Data,
	}, nil
}
