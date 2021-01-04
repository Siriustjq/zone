package db

import (
	"log"
	mydb "zone/src/db/mysql"
)

//UpdateUserFile更新用户文件表
func UpdateUserFile(username, filename, filehash string, filesize int64) bool {
	stmt, err := mydb.DBconnect().Prepare(
		"INSERT INTO tbl_user_file (user_name, file_name, file_sha1, file_size) VALUES (?,?,?,?)")
	if err != nil {
		log.Print(SqlConErr)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, filename, filehash, filesize)
	if err != nil {
		log.Print(SqlExeErr)
		return false
	}
	return true
}
