package handler

import (
	"net/http"

	"pam/greet/internal/logic"
	"pam/greet/internal/svc"
	"pam/greet/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GreetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), ctx) // 接口逻辑处理
		resp, err := l.Greet(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func ByeHandle(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGreetLogic(r.Context(), ctx) // 接口逻辑处理
		resp, err := l.Bye(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
