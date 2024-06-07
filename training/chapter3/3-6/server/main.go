package main

import (
	"context"
	"errors"
	"go-im/training/chapter3/3-6/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServer struct {
}

func (*UserServer) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserResp, error) {
	if u, ok := users[req.Id]; ok {
		return &user.GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}, nil
	}

	return nil, errors.New("不存在查询用户")
}

func main() {
	// 监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("监听失败", err)
	}

	s := grpc.NewServer()
	user.RegisterUserServer(s, &UserServer{})

	log.Println("服务启动成功")

	s.Serve(listener)
}
