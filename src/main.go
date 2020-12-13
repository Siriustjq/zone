package main

import (
	"fmt"
	"log"
	"net/http"
	"zone/src/handler"
)

func main() {
	//设置http的路由规则，类似于Java框架中设置请求拦截规则
	fmt.Printf("1")
	http.HandleFunc("/file/upload", handler.UploadHandler)
	fmt.Printf("2")
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	fmt.Printf("3")
	//开启http监听
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	fmt.Printf("There is an err %s", err.Error())
	//}
	//上面方法不太优雅，现在用log直接包裹监听
	log.Fatal(http.ListenAndServe(":8081", nil))
}
