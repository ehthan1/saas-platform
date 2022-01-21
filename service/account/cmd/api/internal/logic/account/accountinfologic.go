package account

import (
	"context"

	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/svc"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AccountinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) AccountinfoLogic {
	return AccountinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountinfoLogic) Accountinfo() (*types.AccountInfoResponse, error) {
	logx.Infof("UUID: %v", l.ctx.Value("UUID"))
	//
	return &types.AccountInfoResponse{}, nil
}
