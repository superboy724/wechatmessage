package main

import (
	"github.com/superboy724/wechatmessage/processer"
	"github.com/superboy724/wechatmessage/server"
)

func main() {
	server := server.NewServer(80)
	p := processer.NewRegisterProcesser()
	server.SetProcesser(p)
	server.Run()
}
