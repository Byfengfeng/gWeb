package main

import (
	"github.com/Byfengfeng/gWeb/common"
	"github.com/Byfengfeng/gWeb/enum"
	"github.com/Byfengfeng/gWeb/web"
)

func main() {
	webBase := common.NewWebBase()
	webBase.Bind("/login", map[enum.ReqType]func(pkt interface{}) interface{}{
		enum.Post: func(pkt interface{}) interface{} {
			return &User{
				Name: "张飒",
				Age: "20",
			}
		},
		enum.Get: func(pkt interface{}) interface{} {
			return "get"
		},
	})
	web.NewGWebService(":8889",enum.Json,webBase).Start()
}

type User struct {
	Name string `json:"name"`
	Age string `json:"age"`
}