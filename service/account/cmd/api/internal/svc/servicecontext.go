package svc

import (
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/config"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/middleware"
	"github.com/portgas-x/saas-platform/service/account/cmd/rpc/accountclient"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	AccountRpc accountclient.Account
	Jwt        rest.Middleware
	JwtObject  *middleware.JwtMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AccountRpc: accountclient.NewAccount(zrpc.MustNewClient(c.AccountRpc)),
		Jwt:        middleware.NewJwtMiddleware(c).Handle,
	}
}
