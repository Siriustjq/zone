package handler

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
	rp "zone/src/cache/redis"
	dblayer "zone/src/db"
	"zone/src/util"
)

//分块初始化信息
type MultipartUploadInfo struct {
	FileHash   string
	FileSize   int
	UploadID   string
	ChunkSize  int //分块大小
	ChunkCount int //分块数量
}

//初始化分块上传
func InitialMultipartUploadHandler(w http.ResponseWriter, r *http.Request) {
	//1.解析用户参数
	r.ParseForm()
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filesize, err := strconv.Atoi(r.Form.Get("filesize"))
	if err != nil {
		return
	}

	//2.获得一个redis的连接
	rConn := rp.RedisPool().Get()
	defer rConn.Close()

	//3.生成分块上传的初始化信息
	upInfo := MultipartUploadInfo{
		FileHash:   filehash,
		FileSize:   filesize,
		UploadID:   username + fmt.Sprintf("%x", time.Now().UnixNano()),
		ChunkSize:  5 * 1024 * 1024,
		ChunkCount: int(math.Ceil(float64(filesize) / (5 * 1024 * 1024))), //math.Ceil()为向上取整的操作
	}

	//4.将初始化信息写入到redis缓存之中
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "chunkcount", upInfo.ChunkCount) //hash命令：将hash表中的filed域的值设置为value  HSET hash filed value
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "filehash", upInfo.FileHash)
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "filesize", upInfo.FileSize)

	//5.将初始化信息返回到客户端
	w.Write(util.NewRespMsg(0, "OK", upInfo).JSONBytes())
}

//通知上传合并
func CompleteUploadHandler(w http.ResponseWriter, r *http.Request) {
	//1.请求参数
	r.ParseForm()
	upid := r.Form.Get("uploadid")
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filesize := r.Form.Get("filesize")
	filename := r.Form.Get("filename")

	//2.获得redis连接池中的一个连接
	rConn := rp.RedisPool().Get()
	defer rConn.Close()
	//3.通过查询uploadid查询redis并判断是否所有的分块上传完成
	data, err := redis.Values(rConn.Do("HGETALL", "MP_"+upid))
	if err != nil {
		w.Write(util.NewRespMsg(-1, "faild", nil).JSONBytes())
		return
	}
	totalCount := 0
	chunkCount := 0
	for i := 0; i < len(data); i += 2 {
		k := string(data[i].([]byte))
		v := string(data[i+1].([]byte))
		if k == "chunkcoount" {
			totalCount, _ = strconv.Atoi(v)
		} else if strings.HasPrefix(k, "chkid_") && v == "1" {
			chunkCount += 1
		}
	}
	if totalCount != chunkCount {
		w.Write(util.NewRespMsg(-2, "failed", nil).JSONBytes())
		return
	}
	//4.合并分块
	//5.更新唯一文件表和用户文件表
	//todo 和普通文件上传逻辑保持一致
	filesizes, _ := strconv.Atoi(filesize)
	dblayer.SaveFileToDB(filehash, filename, int64(filesizes), "")
	dblayer.UpdateUserFile(username, filename, filehash, int64(filesizes))

	//6.向客户端响应处理结果
	w.Write(util.NewRespMsg(0, "OK", nil).JSONBytes())
}
