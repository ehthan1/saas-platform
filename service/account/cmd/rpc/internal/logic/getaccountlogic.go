package logic

import (
	"context"

	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/account"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *GetAccountLogic) GetAccount(in *account.GetAccountRequest) (*account.GetAccountResponse, error) {
	// todo: add your logic here and delete this line

	return &account.GetAccountResponse{}, nil
}
