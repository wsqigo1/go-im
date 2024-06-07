package main

import (
	"log"
	"net/rpc"
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
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("建立连接失败", err)
	}
	defer client.Close()

	var (
		req  = GetUserReq{Id: "1"}
		resp GetUserResp
	)

	// 调用请求
	err = client.Call("UserServer.GetUser", req, &resp)
	if err != nil {
		log.Println("请求失败", err)
		return
	}

	log.Println(resp)
}
