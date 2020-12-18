package db

import (
	"fmt"
	mydb "zone/src/db/mysql"
)

func SaveFileToDB(filehash string, filename string,
	filesize int64, filelocation string) bool {
	//进行防sql注入操作
	stmt, err := mydb.DBconnect().Prepare(
		"insert ignore into tbl_file('file_sha1','file_name','file_size','file_location','status')" +
			",values(?,?,?,?,1)")
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
