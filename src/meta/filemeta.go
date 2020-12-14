package meta

/**
添加文件元数据类型的结构体FileMeta，方便对文件的操作
*/
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UpdateAt string
}

//目前利用FileSha1建立FileMeta的hashmap索引
var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//UpdataFileMeta新增或者更新fileMetas的map索引
func UpdataFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//GetFileMeta获取文件元信息
func GetFileMeta(fsha1 string) FileMeta {
	return fileMetas[fsha1]
}
