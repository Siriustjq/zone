package main

import (
	"log"
	"net/http"
	"zone/src/handler"
)

func main() {
	//设置http的路由规则，类似于Java框架中设置请求拦截规则
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadFileHandler)
	http.HandleFunc("/file/update", handler.UpdateFileHandler)
	http.HandleFunc("/file/delete", handler.DeleteFileHandler)
	http.HandleFunc("/user/signup", handler.SignUpHandler)
	http.HandleFunc("/user/signin", handler.UserSignInHandler)
	//分块上传接口
	http.HandleFunc("/file/mploaded/init", handler.HttpInterceptor(handler.InitialMultipartUploadHandler))
	http.HandleFunc("file/mploaded/uppart", handler.HttpInterceptor(handler.UploadPartHandler))
	http.HandleFunc("file/mploaded/complete", handler.HttpInterceptor(handler.CompleteUploadHandler))
	//添加拦截器
	http.HandleFunc("/user/info", handler.HttpInterceptor(handler.UserInfoHandler))
	//统筹静态资源
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	//开启http监听
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil
	//	fmt.Printf("There is an err %s", err.Error())
	//}
	//上面方法不太优雅，现在用log直接包裹监听
	log.Fatal(http.ListenAndServe(":8081", nil))
}
