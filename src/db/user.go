package db

import (
	"fmt"
	"log"
	mydb "zone/src/db/mysql"
)

//规范一些日志输出字段
const (
	SqlConErr = "It's wrong when try to connect mysql"
	SqlExeErr = "It's wrong when try to get data from mysql"
)

//UserSingnUp用户注册接口设计
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

var name = "default"
var pass = "123456"

//UserSignIn用户登录接口设计
func UserSignIn(username string, password string) bool {
	stmt, err := mydb.DBconnect().Prepare("SELECT user_name, user_pwd FROM tbl_user WHERE user_name = ? LIMIT 1")
	if err != nil {
		log.Print(SqlConErr)
		return false
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&name, &pass)
	if err != nil {
		log.Print(SqlExeErr)
		return false
	}
	fmt.Print(name + " " + password)
	if name == "" {
		return false
	} else if pass != password {
		return false
	}
	return true
}

func UpdateUserToken(username string, token string) bool {
	stmt, err := mydb.DBconnect().Prepare("REPLACE INTO tbl_user_token (user_name,user_token) VALUES (?,?)")
	if err != nil {
		log.Print(SqlConErr)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, token)
	if err != nil {
		log.Print(SqlExeErr)
		return false
	}
	return true
}

//根据username查询用户信息
type UserInfo struct {
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

func GetUserInfo(username string) (UserInfo, error) {
	userinfo := UserInfo{}
	stmt, err := mydb.DBconnect().Prepare("SELECT user_name, signup_at FROM tbl_user WHERE user_name = ? LIMIT 1")
	if err != nil {
		log.Print(SqlConErr)
		return UserInfo{}, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&userinfo.Username, &userinfo.SignupAt)
	if err != nil {
		log.Print(SqlExeErr)
		return UserInfo{}, err
	}
	return userinfo, nil
}
