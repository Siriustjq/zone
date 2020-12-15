package handler

/**
实现文件的上传和下载
*/
import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"zone/src/meta"
	"zone/src/util"
)

//文件上传(这里一定要注意，方法名首字母大写，否则无法在别的包中被引用发现)
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//首次访问指定url默认采用GET方法提交，所以需要调出提交文件表单页面
	if r.Method == "GET" {
		//通过读取html文件再交由http.ResponseWriter输出的方式实现文件提交页面的唤出
		data, err := ioutil.ReadFile("static/view/index.html")
		if err != nil {
			_, _ = io.WriteString(w, "something wrong!")
			return
		}
		_, _ = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//将文件存储至本地
		file, head, err := r.FormFile("file")

		if err != nil {
			fmt.Printf("Failed to get file data %s\n", err.Error())
			return
		}
		defer file.Close()
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		//在本地创建一个新的文件去承载上传的文件
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create newFile data %s\n", err.Error())
			return
		}

		defer newFile.Close()
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save into newFile %s\n", err.Error())
			return
		}
		//将文件光标移至文件开头，且偏移量为0
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		fmt.Printf(fileMeta.FileSha1)
		//将刚刚上传的文件的sha1索引添加到map中
		meta.UpdataFileMeta(fileMeta)
		// 重定向到成功的页面逻辑
		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

// 文件上传成功处理逻辑
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Upload Succeed!")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	//格式化请求参数信息
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	//将文件转换为json格式
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//仍然是通过前面定义的sha1串进行唯一索引
	filehash := r.Form.Get("filehash")
	filemeta := meta.GetFileMeta(filehash)
	file, err := os.Open(filemeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//为在浏览器中演示，添加http头，让浏览器弹出下载页面
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-disposition", "attachment;filename=\""+filemeta.FileName+"\"")
	w.Write(data)
}
