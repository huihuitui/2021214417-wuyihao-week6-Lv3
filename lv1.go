package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func initmysql() (err error) {
	count = 1
	Usermap = make(map[string]string)
	var dns = "root:123456@(localhost:3306)/db3?parseTime=true&loc=Local" //DSN（数据源名称）
	db, err = sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(2)
	return err
}
func insertsql() {
	rsl, err := db.Exec("insert into user(name,password,passquestion,passquestion_ans) value(?,?,?,?)", user, password, qus, ans)
	if err != nil {
		log.Println(err)
	}
	n, _ := rsl.RowsAffected()
	fmt.Println(n)
}
func Ostin() {
	fmt.Println("输入用户名")
	fmt.Scan(&user)
	fmt.Println("输入密码")
	fmt.Scan(&password)
	fmt.Println("输入密保问题")
	fmt.Scan(&qus)
	fmt.Println("输入密保")
	fmt.Scan(&ans)
	Usermap[user] = password
	fmt.Println("ok")
}
func Middlewar() gin.HandlerFunc {
	var User struct {
		Name     string
		Password string
	}
	return func(c *gin.Context) {
		choice := getlist([]string{"登录用户", "修改密码"})
		switch choice {
		case 1:
			fmt.Println("输入用户名")
			fmt.Scan(&user)
			fmt.Println("输入密码")
			fmt.Scan(&password)
			row := db.QueryRow("select name,password from user where name =?", user)
			row.Scan(&User.Name, &User.Password)
			if User.Password == password && User.Name == user {
				c.Next()
				return
			}
			c.String(400, "未注册")
			c.Abort()
			return
		case 2:
			UpdatePassword()
		default:
			fmt.Println("没有此选项")
		}
	}
}
