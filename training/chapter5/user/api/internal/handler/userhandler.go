package handler

import (
	"net/http"

	"github.com/wsqigo/training/docker/user/api/internal/logic"
	"github.com/wsqigo/training/docker/user/api/internal/svc"
	"github.com/wsqigo/training/docker/user/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.User(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}