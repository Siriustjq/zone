package db

import (
	"database/sql"
	"fmt"
	"log"
	mydb "zone/src/db/mysql"
)

//插入数据
func SaveFileToDB(filehash string, filename string,
	filesize int64, filelocation string) bool {
	//进行防sql注入操作
	stmt, err := mydb.DBconnect().Prepare(
		"INSERT INTO tbl_file (file_sha1,file_name,file_size,file_addr,status)VALUES(?,?,?,?,1)")
	if err != nil {
		fmt.Println("Failed to prepare statement, the err is ", err.Error())
		return false
	}
	//语句对象一定要关掉
	defer stmt.Close()
	//执行语句
	ret, err := stmt.Exec(filehash, filename, filesize, filelocation)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	//判断是否是更新
	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Println("This file already exits, just update right now")
		}
		return true
	}
	//err != nil 插入失败
	return false
}

//获取数据
type TableFile struct {
	FileHash sql.NullString
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//GetFileFromDB从数据库中获取文件信息
func GetFileFromDB(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBconnect().Prepare(
		"SELECT file_sha1,file_name,file_size,file_addr FROM tbl_file WHERE file_sha1 = ? AND status = '1' LIMIT 1")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()
	res := TableFile{}
	//按行发出请求，并格式化获取返回值
	err = stmt.QueryRow(filehash).Scan(&res.FileHash, &res.FileName,
		&res.FileSize, &res.FileAddr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &res, nil
}
