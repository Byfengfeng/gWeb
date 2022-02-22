package web

import (
	"encoding/json"
	"github.com/Byfengfeng/gWeb/common"
	"github.com/Byfengfeng/gWeb/enum"
	"github.com/Byfengfeng/gWeb/inter"
	"net/http"
)




type gWebService struct {
	addr string
	contentType enum.ContentType
	reqFn common.WebBase
}

func NewGWebService(addr string,contentType enum.ContentType,ReqFn common.WebBase) inter.IGWeb {
	return &gWebService{addr: addr,contentType: contentType,reqFn: ReqFn}
}


func (g *gWebService) Start()  {
	for url,reqMap := range g.reqFn.ReqMap {
		http.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Set("Cache-Control","no-cache")
			writer.Header().Set("Access-Control-Allow-Origin", "*")
			writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			writer.Header().Set("Access-Control-Allow-Credentials", "true")
			var res interface{}
			method := request.Method

			if method == enum.Get.GetString(){
				hanDel,ok := reqMap[enum.Get]
				if ok {
					values := request.URL.Query()
					if len(values) > 0 {
						res = hanDel(values)
					}
				}
			}else if method == enum.Post.GetString(){
				hanDel,ok := reqMap[enum.Post]
				if ok {
					request.FormValue("")
					if g.contentType == enum.From {
						if request.Form != nil {
							if len(request.Form) > 0 {
								token := request.Header.Get("token")
								if token != "" {
									request.Form["token"] = []string{token}
								}

								res = hanDel(request.Form)
							}
						}
					} else if g.contentType == enum.Json{
						if request.PostForm != nil {
							if len(request.PostForm) > 0 {
								res = hanDel(request.PostForm)
							}
						}
					}
				}
			}

			if res != nil {
				writer.Header().Set("Content-Type", "application/json; charset=utf-8")
				bytes, err := json.Marshal(&res)
				if err == nil {
					writer.Write(bytes)
				}
			}

		})
	}
	err := http.ListenAndServe(g.addr, nil)
	if err != nil {
		panic(err)
	}
}