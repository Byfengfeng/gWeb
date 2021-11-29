package main

import (
	"github.com/Byfengfeng/gWeb/common"
	"github.com/Byfengfeng/gWeb/enum"
	"github.com/Byfengfeng/gWeb/util"
	"github.com/Byfengfeng/gWeb/web"
)

func main() {
	webBase := common.NewWebBase()
	webBase.Bind("/login", map[enum.ReqType]func(pkt map[string][]string) interface{}{
		enum.Post: func(pkt map[string][]string) interface{} {
			return &User{
				Name: "张飒",
				Age: 20,
			}
		},
		enum.Get: func(pkt map[string][]string) interface{} {
			req := util.UnmarshalReq(pkt,&User{}).(*User)
			req.Age = 20
			return req
		},
	})
	web.NewGWebService(":8889",enum.Json,webBase).Start()
}

type User struct {
	Name string `json:"name"`
	Age int32 `json:"age"`
}