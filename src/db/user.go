package db

import (
	"fmt"
	"log"
	mydb "zone/src/db/mysql"
)

func UserSingnUp(username string, password string) bool {
	stmt, err := mydb.DBconnect().Prepare(
		"INSERT INTO tbl_user (user_name,user_pwd,status)VALUES(?,?,1)")
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	defer stmt.Close()
	res, err := stmt.Exec(username, password)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	//判断是否是更新
	if rf, err := res.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Println("This file already exits, just update right now")
		}
		return true
	}
	//err != nil 插入失败
	return false
}
