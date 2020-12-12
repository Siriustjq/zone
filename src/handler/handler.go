package handler

/**
实现文件的上传和下载
*/
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//文件上传
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	//默认采用post方法提交，所以如果是get则重新返回值文件上传界面
	if r.Method == "GET" {

	} else if r.Method == "Post" {
		//将文件存储至本地
	}
}
