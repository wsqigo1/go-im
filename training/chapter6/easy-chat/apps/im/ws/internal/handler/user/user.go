package user

import (
	"github.com/wsqigo/easy-chat/apps/im/ws/internal/svc"
	"github.com/wsqigo/easy-chat/apps/im/ws/websocket"
)

func OnLine(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		uids := srv.GetUsers()
		u := srv.GetUsers(conn)
		err := srv.Send(websocket.NewMessage(u[0], uids), conn)
		srv.Info("err", err)
	}
}
