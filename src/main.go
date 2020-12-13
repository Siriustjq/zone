package main

import (
	"fmt"
	"net/http"
	"zone/src/handler"
)

func main() {
	//设置http的路由规则，类似于Java框架中设置请求拦截规则
	http.HandleFunc("/file/upload", handler.UploadHandler)
	//开启http监听
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("There is an err %s", err.Error())
	}
}
