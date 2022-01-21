package account

import (
	"context"

	"github.com/portgas-x/saas-platform/common/utils"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/svc"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) RefreshTokenLogic {
	return RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken() (*types.JwtToken, error) {
	token, err := utils.BuildToken(l.svcCtx.Config.Auth.AccessSecret, map[string]interface{}{
		"UUID": l.ctx.Value("UUID"),
	}, int64(l.svcCtx.Config.Auth.AccessExpire))
	if err != nil {
		return nil, err
	}
	return &types.JwtToken{
		AccessToken:  token,
		AccessExpire: int64(l.svcCtx.Config.Auth.AccessExpire),
	}, nil
}
