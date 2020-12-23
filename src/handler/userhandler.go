package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	mydb "zone/src/db"
	"zone/src/util"
)

const ex = "tjq"

/**
主要来处理用户注册的一些业务逻辑
*/
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	//和之前逻辑一样，默认发起的GET请求，需要通过GET请求将注册页面吐出去
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("static/view/signup.html")
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		_, _ = w.Write(data)
		//或者通过下面方法返回该页面
		//_ , _ = io.WriteString(w,string(data))
	} else {
		_ = r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		//为密码加密
		if len(username) > 10 || len(password) < 5 {
			_, _ = w.Write([]byte("Invalid Parameters!"))
			log.Print("Invalid Parameters!")
			log.Fatal()
			return
		}
		password = util.Sha1([]byte(password + ex))
		res := mydb.UserSingnUp(username, password)
		if res {
			_, _ = w.Write([]byte("SUCCESS"))
			log.Print("SUCCESS")
			return
		} else {
			_, _ = w.Write([]byte("FAILED"))
			log.Print("FAILED")
			return
		}
	}
}

/**
用户登录模块接口设计
*/
func UserSignInHandler(w http.ResponseWriter, r *http.Request) {
	//和之前逻辑一样，默认发起的GET请求，需要通过GET请求将注册页面吐出去
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("static/view/signin.html")
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		_, _ = w.Write(data)
		//或者通过下面方法返回该页面
		//_ , _ = io.WriteString(w,string(data))
	} else {
		//从数据库中拉取查询该登录用户是否存在
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		password = util.Sha1([]byte(password + ex))
		if mydb.UserSignIn(username, password) {
			//添加用户认证token
			token := GenerateToken(username)
			//向数据库中更新用户的token信息
			if mydb.UpdateUserToken(username, token) {
				//成功重定向到主页
				w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
			} else {
				w.Write([]byte("更新用户token失败"))
				return
			}
		} else {
			w.Write([]byte("该用户未注册"))
			return
		}
	}
}

//生成用户token
func GenerateToken(name string) string {
	//取name+当前时间+ex的md5值再加前八位的当前时间
	t := fmt.Sprintf("%x", time.Now().Unix())
	return util.MD5([]byte(name+t+ex)) + t[:8]
}
