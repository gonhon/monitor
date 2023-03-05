package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"monitor/monitor/internal/logic"
	"monitor/monitor/internal/svc"
)

func getMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetMsgLogic(r.Context(), svcCtx)
		resp, err := l.GetMsg()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
