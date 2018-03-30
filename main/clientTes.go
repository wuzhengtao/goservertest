package main

import (
	"goservertest/login"
	"encoding/json"
	"goservertest/define"
	"net"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

//客户端注册测试

func main() {
	//signuptes()测试成功
	//signintes()测试成功
	//updatetes()测试成功
}

//注册测试
func signuptes() {
	signupmsg := login.SignUpMsg{"client1@163.com", "123456"}
	smbyte, err := json.Marshal(signupmsg)
	checkError(err)

	code := int(101)
	msg := define.Message{code, smbyte}
	msgbyte, err := json.Marshal(msg)
	checkError(err)

	conn, err := net.Dial(define.TCP, define.HOST)
	buf := make([]byte, 50)
	checkError(err)
	defer conn.Close()
	//go heartMinute(conn)

	conn.Write(msgbyte)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("send msg failed")
	}
	receive := buf[0:n]
	recMsg := &define.Message{}
	json.Unmarshal(receive, recMsg)
	fmt.Println("answer is ", recMsg.Code)
	fmt.Println("send msg successfully")

	fmt.Println("send msg end")
}

//登录测试
func signintes() {
	signinmsg := login.SignInMsg{SignType: "Email", Value: "client1@163.com", Password: "123456"}
	smbyte, err := json.Marshal(signinmsg)
	checkError(err)

	code := int(111)
	msg := define.Message{code, smbyte}
	msgbyte, err := json.Marshal(msg)
	checkError(err)

	conn, err := net.Dial(define.TCP, define.HOST)
	buf := make([]byte, 50)
	checkError(err)
	defer conn.Close()
	//go heartMinute(conn)

	conn.Write(msgbyte)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("send msg failed")
	}
	receive := buf[0:n]
	recMsg := &define.Message{}
	json.Unmarshal(receive, recMsg)
	fmt.Println("answer is ", recMsg.Code)
	fmt.Println("send msg successfully")

	fmt.Println("send msg end")
}

//更新测试
func updatetes() {
	updatemsg := login.UpdateMsg{UserName: "client6", Telephone: "12345678910", Email: "123123123@163.com", Sex: 1, Id: 6}
	smbyte, err := json.Marshal(updatemsg)
	checkError(err)

	code := int(121)
	msg := define.Message{Code: code, Msg: smbyte}
	msgbyte, err := json.Marshal(msg)
	checkError(err)

	conn, err := net.Dial(define.TCP, define.HOST)
	buf := make([]byte, 50)
	checkError(err)
	defer conn.Close()
	//go heartMinute(conn)

	conn.Write(msgbyte)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("send msg failed")
	}
	receive := buf[0:n]
	recMsg := &define.Message{}
	json.Unmarshal(receive, recMsg)
	fmt.Println("answer is ", recMsg.Code)
	fmt.Println("send msg successfully")

	fmt.Println("send msg end")
}

//密码更新测试
func updatepasswordtes() {
	updatemsg := login.UpdatePW{Password: "12345678910", Id: 6}
	smbyte, err := json.Marshal(updatemsg)
	checkError(err)

	code := int(131)
	msg := define.Message{Code: code, Msg: smbyte}
	msgbyte, err := json.Marshal(msg)
	checkError(err)

	conn, err := net.Dial(define.TCP, define.HOST)
	buf := make([]byte, 50)
	checkError(err)
	defer conn.Close()
	//go heartMinute(conn)

	conn.Write(msgbyte)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("send msg failed")
	}
	receive := buf[0:n]
	recMsg := &define.Message{}
	json.Unmarshal(receive, recMsg)
	fmt.Println("answer is ", recMsg.Code)
	fmt.Println("send msg successfully")

	fmt.Println("send msg end")
}
