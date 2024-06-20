package push

import (
	"github.com/wsqigo/easy-chat/apps/im/ws/internal/svc"
	"github.com/wsqigo/easy-chat/apps/im/ws/websocket"
)

func Push(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {

	}
}
