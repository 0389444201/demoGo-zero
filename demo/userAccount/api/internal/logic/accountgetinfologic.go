package logic

import (
	"context"
	"fmt"

	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountGetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountGetInfoLogic {
	return &AccountGetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountGetInfoLogic) AccountGetInfo(req *types.InfoRequest) (resp *types.InfoResponse, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("Process login account &v", req.Name)
	fmt.Println(req.Name)
	info, err := l.svcCtx.UserModel.FindByName(l.ctx, req.Name)
	return &types.InfoResponse{
		Name:  info.Name,
		Email: info.Email,
	}, nil
}
