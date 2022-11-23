package main

import (
	"fmt"
	"log"
)

func UpdatePassword() {
	var ans string
	var User struct {
		name             string
		password         string
		passwordquestion string
		questionAns      string
	}
	fmt.Println("输入需要修改密码的账户")
	fmt.Scan(&user)
	rows := db.QueryRow("select * from user where name=?", user)
	err := rows.Scan(&User.name, &User.password, &User.passwordquestion, &User.questionAns)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(User.passwordquestion)
	fmt.Println("请输入答案")
	fmt.Scan(&ans)
	if ans == User.questionAns {
		fmt.Println("请输入新密码")
		fmt.Scan(&password)
		Usermap[user] = password
		rsl, err := db.Exec("update user set password=? where name=?", password, user)
		if err != nil {
			log.Println(err)
		}
		n, _ := rsl.RowsAffected()
		fmt.Println(n)
		return
	}
	fmt.Println("输入错误")
	return
}
