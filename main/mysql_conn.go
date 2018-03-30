package main

import (
	_ "github.com/go-sql-driver/mysql"
	"goservertest/login"
	"fmt"
	"goservertest/db"
	"goservertest/define"
)

func main() {
	var dbp = db.NewProcessor()
	var lp = login.NewProcessor(dbp)
	user := &define.User{}
	user.Sex = define.MALE
	user.UserName = "tyy"
	user.Telephone = "15555555"
	user.Email = "jll@126.com"
	if err:=lp.SignUp(user.Email, "987654321"); err!=nil{
		fmt.Println(err)
	}
}
