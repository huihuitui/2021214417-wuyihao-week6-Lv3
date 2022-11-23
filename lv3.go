package main

import (
	"fmt"
	"log"
	"time"
)

func Newcomment() {
	var message string
	fmt.Println("输入评论内容")
	fmt.Scan(&message)
	rsl, err := db.Exec("insert into message_board (user_id,time,message,message_id) values (?,?,?,?)", Usermap["user"], time.Now(), message, count)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rsl.RowsAffected())
	count++
}
func Recomment() {
	var message string
	var id uint
	fmt.Println("输入评论内容")
	fmt.Scan(&message)
	fmt.Println("输入回复的评论id")
	fmt.Scan(&id)
	rsl, err := db.Exec("insert into message_board (user_id,time,message,message_id) values (?,?,?,?)", Usermap["user"], time.Now(), message, id)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rsl.RowsAffected())
}
func showmessage() {
	var message struct {
		name      string
		time      time.Time
		value     string
		messageid uint
	}
	rows, err := db.Query("select * from message_board")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&message.name, &message.time, &message.value, &message.messageid)
		fmt.Printf("time:%v", message.time)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%+v\t\n", message)
	}
}
