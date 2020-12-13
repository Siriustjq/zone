package handler

/**
实现文件的上传和下载
*/
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//文件上传(这里一定要注意，方法名首字母大写，否则无法在别的包中被引用发现)
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("4")
	//首次访问指定url默认采用GET方法提交，所以需要调出提交文件表单页面
	if r.Method == "GET" {
		fmt.Printf("5")
		//通过读取html文件再交由http.ResponseWriter输出的方式实现文件提交页面的唤出
		data, err := ioutil.ReadFile("static/view/index.html")
		if err != nil {
			_, _ = io.WriteString(w, "something wrong!")
			return
		}
		_, _ = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		fmt.Printf("6")
		//将文件存储至本地
		file, head, err := r.FormFile("file")

		if err != nil {
			fmt.Printf("Failed to get file data %s\n", err.Error())
			return
		}
		defer file.Close()
		//在本地创建一个新的文件去承载上传的文件
		newFile, err := os.Create("/tmp/" + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create newFile data %s\n", err.Error())
			return
		}

		defer newFile.Close()
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save into newFile %s\n", err.Error())
			return
		}
		// 重定向到成功的页面逻辑
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// 文件上传成功处理逻辑
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Upload Succeed!")
}
