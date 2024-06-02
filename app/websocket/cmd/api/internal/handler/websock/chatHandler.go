package websock

import (
	"golodge/app/websocket/cmd/api/internal/logic/process"
	"golodge/app/websocket/cmd/api/internal/logic/websock"
	"golodge/app/websocket/cmd/api/internal/svc"
	"net/http"
)

// chat
func ChatHandler(svcCtx *svc.ServiceContext, hub *process.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := websock.NewChatLogic(r.Context(), svcCtx)
		l.Chat(w, r, hub)
	}
}
