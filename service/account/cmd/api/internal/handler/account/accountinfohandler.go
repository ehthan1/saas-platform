package account

import (
	"net/http"

	"github.com/portgas-x/saas-platform/common/response"
	account "github.com/portgas-x/saas-platform/service/account/cmd/api/internal/logic/account"
	"github.com/portgas-x/saas-platform/service/account/cmd/api/internal/svc"
)

func AccountinfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewAccountinfoLogic(r.Context(), svcCtx)
		resp, err := l.Accountinfo()
		response.Response(w, resp, err)
	}
}
