package handler

/**
实现文件的上传和下载
*/
import (
	"io"
	"io/ioutil"
	"net/http"
)

//文件上传(这里一定要注意，方法名首字母大写，否则无法在别的包中被引用发现)
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//默认采用post方法提交，所以如果是get则重新返回值文件上传界面
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("static/view/home.html")
		if err != nil {
			io.WriteString(w, "something wrong!")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "Post" {
		//将文件存储至本地
	}
}
