package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

/**
对mysql数据库进行连接，database/sql是go语言连接数据库的标准接口，
"github.com/go-sql-driver/mysql"则是go语言连接mysql的驱动
*/

//创建数据库对象
var db *sql.DB

//初始化
func init() {
	db, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/quming?charset=utf8")
	//设置最大连接数
	db.SetMaxOpenConns(1000)
	//测试连接的通断情况
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed connect to mysql", err.Error())
		//强制当前线程退出
		os.Exit(1)
	}
}

//提供外部接口，返回数据库连接对象
func DBconnect() *sql.DB {
	return db
}
