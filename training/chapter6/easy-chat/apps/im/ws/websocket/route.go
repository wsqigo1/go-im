package websocket

type Route struct {
	Method  string
	Handler HandlerFunc
}

// conn 当前请求的用户的连接
type HandlerFunc func(srv *Server, conn *Conn, msg *Message)
