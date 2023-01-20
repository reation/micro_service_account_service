package logic

import (
	"context"
	"fmt"
	"github.com/reation/micro_service_account_service/account/internal/svc"
	"github.com/reation/micro_service_account_service/config"
	"github.com/reation/micro_service_account_service/model"
	"github.com/reation/micro_service_account_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountLogic {
	return &GetAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountLogic) GetAccount(in *protoc.GetAccountRequest) (*protoc.GetAccountResponse, error) {
	accountInfo, err := l.svcCtx.AccountModel.GetAccountInfoByUserID(l.ctx, in.GetUserId())
	if err == model.ErrNotFound {
		return &protoc.GetAccountResponse{States: config.RETURN_STATE_USER_NOT_FOUND, UserId: in.GetUserId(), Balance: 0.00}, nil
	}
	fmt.Println("-------------------------------")
	fmt.Println(accountInfo)
	fmt.Println("-------------------------------")
	if accountInfo.State == model.STATE_UNUSUAL {
		return &protoc.GetAccountResponse{States: config.RETURN_STATE_USER_UNUSUAL, UserId: in.GetUserId(), Balance: accountInfo.Balance}, nil
	}

	if accountInfo.State == model.STATE_FORZEN {
		return &protoc.GetAccountResponse{States: config.RETURN_STATE_USER_FORZEN, UserId: in.GetUserId(), Balance: accountInfo.Balance}, nil
	}

	return &protoc.GetAccountResponse{States: config.RETURN_STATE_NORMAL, UserId: in.GetUserId(), Balance: accountInfo.Balance}, nil
}
