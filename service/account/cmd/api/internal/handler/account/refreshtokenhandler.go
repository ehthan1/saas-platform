package account

import (
	"net/http"

	"github.com/portgas-x/saas-platform/common/response"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/logic/account"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/svc"
)

func RefreshTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewRefreshTokenLogic(r.Context(), svcCtx)
		resp, err := l.RefreshToken()
		response.Response(w, resp, err)
	}
}
