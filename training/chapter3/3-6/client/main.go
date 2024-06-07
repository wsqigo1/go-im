package main

import (
	"context"
	"go-im/training/chapter3/3-6/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GetUserReq struct {
	Id string `json:"id"`
}

type GetUserResp struct {
	Id    string
	Name  string
	Phone string
}

func main() {
	client, err := grpc.NewClient("localhost:1234",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("建立连接失败", err)
	}
	defer client.Close()

	c := user.NewUserClient(client)
	resp, err := c.GetUser(context.Background(), &user.GetUserReq{
		Id: "1",
	})
	if err != nil {
		log.Fatal("请求失败", err)
	}

	log.Println(resp)
}
