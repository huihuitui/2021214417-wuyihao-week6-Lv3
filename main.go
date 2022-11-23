package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var user, password, qus, ans string
var db *sql.DB
var Usermap map[string]string
var count uint

func main() {
	err := initmysql()
	if err != nil {
		fmt.Println("连接失败")
	}
	defer db.Close()
	r := gin.Default()
	r.GET("/register", func(c *gin.Context) {
		Ostin()
		insertsql()
		Usermap["user"] = user
		c.SetCookie(user, password, 120, "/", "localhost", false, true)
		c.String(200, "已经自动登录")
	})
	r.GET("/login", Middlewar(), func(c *gin.Context) {
		Usermap["user"] = user
		c.String(200, "登陆成功")
		for {
			choice := getlist([]string{"新建留言", "对留言回复", "查看留言板", "退出"})
			switch choice {
			case 1:
				Newcomment()
			case 2:
				Recomment()
			case 3:
				showmessage()
			case 4:
				return
			default:
				fmt.Println("没有此选项")
			}
		}
	})

	r.Run()
}
func getlist(slist []string) int {
	fmt.Println("***************选择栏**************")
	for i, v := range slist {
		fmt.Printf("%v:%s", i+1, v)
	}
	for {
		var choice int
		fmt.Scan(&choice)
		if choice <= 0 || choice > len(slist) {
			continue
		}
		return choice
	}
}
