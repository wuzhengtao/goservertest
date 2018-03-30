package login

import (
	"log"
	"encoding/json"
	"goservertest/define"
	"fmt"
)

var lp *Processor
var ch chan []uint8

type Proxy struct {
}

/**
proxy类主要负责接收对应逻辑的信息，然后处理逻辑

ProcessMessage ：
	proxy 类中唯一暴露出来的方法。接收信息，处理信息，返回信息
checkCode :
	这个方法检查传过来的代码有没有错误
	login中主要的消息有6个，分别对应101～151，检查首位与末位应该都为1

对各种信息处理的方法，输入：[]uint8， 输出[]uint8
 */
func (p *Proxy) ProcessMessage(code int, msg []uint8) {
	fmt.Println("sProcessMessagetart")
	if !checkCode(code) {
		log.Print("code is error")
		return
	}

	var ans []uint8

	switch code {
	case 101:
		//注册，暂定邮箱登录，之后可以更新手机号等信息，
		ans = signup(code, msg)
	case 111:
		//登录，登录支持邮箱登录，手机登录，用户名登录
		ans = signin(code, msg)
	case 121:
		//更新用户信息，主要修改四个信息，用户名、手机号、邮箱、性别
		ans = update(code, msg)
	case 131:
		//更新密码
		ans = update_password(code, msg)
	case 141:
		//退出账号
		//ans = signout(code, msg)
	case 151:
		//注销账号
		//ans = logout(code, msg)
	}

	ch <- ans
}

func update(code int, msg []uint8) []uint8 {
	requestmsg := &UpdateMsg{}
	json.Unmarshal(msg, requestmsg)
	id := requestmsg.Id
	username := requestmsg.UserName
	telephone := requestmsg.Telephone
	email := requestmsg.Email
	sex := requestmsg.Sex

	User := &define.User{UserName: username, Telephone: telephone, Email: email, Sex: sex}

	err := lp.Update(User, id)

	if err != nil {
		code = code + 2
		errmsg := define.ErrorMsg{Code: code, Err: err.Error()}
		ans, _ := json.Marshal(errmsg)
		return ans
	}

	code = code + 1
	message := make([]uint8, 0)
	normalmsg := define.Message{Code: code, Msg: message}

	ans, _ := json.Marshal(normalmsg)
	return ans

}

func update_password(code int, msg []uint8) []uint8 {
	requestmsg := &UpdatePW{}
	json.Unmarshal(msg, requestmsg)
	id := requestmsg.Id
	password := requestmsg.Password

	err := lp.UpdatePassword(id, password)

	if err != nil {
		code = code + 2
		errmsg := define.ErrorMsg{Code: code, Err: err.Error()}
		ans, _ := json.Marshal(errmsg)
		return ans
	}

	code = code + 1
	message := make([]uint8, 0)
	normalmsg := define.Message{Code: code, Msg: message}

	ans, _ := json.Marshal(normalmsg)
	return ans
}
//
//func signout(code int, msg []uint8) []uint8 {
//
//}

//func logout(code int, msg []uint8) []uint8 {
//
//}

func signup(code int, msg []uint8) []uint8 {
	requestmsg := &SignUpMsg{}
	json.Unmarshal(msg, requestmsg)
	email := requestmsg.Email
	password := requestmsg.Password

	err := lp.SignUp(email, password)
	if err != nil {
		code = code + 2
		errmsg := define.ErrorMsg{Code: code, Err: err.Error()}
		ans, _ := json.Marshal(errmsg)
		return ans
	}

	code = code + 1
	message := make([]uint8, 0)
	normalmsg := define.Message{Code: code, Msg: message}

	ans, _ := json.Marshal(normalmsg)
	return ans
}

func signin(code int, msg []uint8) []uint8 {
	requestmsg := &SignInMsg{}
	json.Unmarshal(msg, requestmsg)
	signtype := requestmsg.SignType
	value := requestmsg.Value
	password := requestmsg.Password

	id, err := lp.SignIn(signtype, value, password)
	if err != nil {
		code = code + 2
		errmsg := define.ErrorMsg{Code: code, Err: err.Error()}
		ans, _ := json.Marshal(errmsg)
		return ans
	}

	code = code + 1
	message := make([]uint8, 1)
	message[0] = uint8(id)
	normalmsg := define.Message{Code: code, Msg: message}

	ans, _ := json.Marshal(normalmsg)
	return ans
}

func checkCode(code int) bool {
	return code/100 == 1 && code%10 == 1
}

func NewProxy(processor *Processor, c chan []uint8) *Proxy {
	proxy := &Proxy{}
	lp = processor
	ch = c
	return proxy
}
