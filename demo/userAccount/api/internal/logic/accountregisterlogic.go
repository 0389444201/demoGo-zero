package logic

import (
	"context"
	"demo/userAccount/api/internal/helpers"
	"demo/userAccount/api/internal/svc"
	"demo/userAccount/api/internal/types"
	"demo/userAccount/models"
	"github.com/zeromicro/go-zero/core/logx"
)

type AccountRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountRegisterLogic {
	return &AccountRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountRegisterLogic) AccountRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	l.Logger.Info("Process register %v", req)
	user := models.UserTable{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Gender:   req.Gender,
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &user)
	password := helpers.HashPassword(req.Password)
	req.Password = password
	if err != nil {
		l.Logger.Errorf("failed to register account")
	}
	return &types.RegisterResponse{
		Name:   req.Name,
		Email:  req.Email,
		Gender: req.Gender,
	}, nil
}
