package logic

import (
	"context"

	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/account"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAccountLogic {
	return &CreateAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAccountLogic) CreateAccount(in *account.CreateAccountRequest) (*account.CreateAccountResponse, error) {
	// todo: add your logic here and delete this line

	return &account.CreateAccountResponse{}, nil
}
