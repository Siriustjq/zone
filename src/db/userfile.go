package db

import (
	"log"
	"time"
	mydb "zone/src/db/mysql"
)

//构建用户文件结构体，用于用户文件查询的数据返回
type UserFile struct {
	UserName    string
	FileName    string
	FileHash    string
	FileSize    int64
	UploadAt    string
	LastUpdated string
}

//UpdateUserFile更新用户文件表
func UpdateUserFile(username, filename, filehash string, filesize int64) bool {
	stmt, err := mydb.DBconnect().Prepare(
		"INSERT INTO tbl_user_file (user_name, file_name, file_sha1, file_size) VALUES (?,?,?,?)")
	if err != nil {
		log.Print(SqlConErr)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, filename, filehash, filesize, time.Now())
	if err != nil {
		log.Print(SqlExeErr)
		return false
	}
	return true
}

//QueryUserFile通过username获取指定limit行数的用户文件信息
func QueryUserFile(username string, limit int64) ([]UserFile, error) {
	stmt, err := mydb.DBconnect().Prepare("SELECT file_name, file_sha1, file_size, upload_at, last_update " +
		"FROM tbl_user_file WHERE user_name = ? LIMIT ?")
	if err != nil {
		log.Print(SqlConErr)
		return nil, err
	}
	defer stmt.Close()
	row, err := stmt.Query(username, limit)
	if err != nil {
		log.Print(SqlExeErr)
		return nil, err
	}
	var userfiles []UserFile
	for row.Next() {
		ufile := UserFile{}
		err = row.Scan(&ufile.FileName, &ufile.FileHash, &ufile.FileSize, &ufile.UploadAt, &ufile.LastUpdated)
		if err != nil {
			log.Print(SqlExeErr)
			break
		}
		userfiles = append(userfiles, ufile)
	}
	return userfiles, nil
}
