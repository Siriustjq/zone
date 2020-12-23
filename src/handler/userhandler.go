package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	mydb "zone/src/db"
	"zone/src/util"
)

/**
主要来处理用户登录等相关的一些业务逻辑
*/
const ex = "tjq"

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
