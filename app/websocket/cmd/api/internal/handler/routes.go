// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"golodge/app/websocket/cmd/api/internal/logic/process"
	"net/http"

	websock "golodge/app/websocket/cmd/api/internal/handler/websock"
	"golodge/app/websocket/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext, hub *process.Hub) {
	server.AddRoutes(
		[]rest.Route{
			{
				// chat
				Method:  http.MethodGet,
				Path:    "/chat",
				Handler: websock.ChatHandler(serverCtx, hub),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/websocket/v1"),
	)
}
