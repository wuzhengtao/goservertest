package login

import (
	"goservertest/db"
	"goservertest/define"
	"errors"
	"log"
)

var dbp *db.Processor

type Processor struct {
}

//注册，主要验证用户名有没有重复，或者说邮箱或手机有没有注册过
func (p *Processor) SignUp(Email string, Password string) error {
	var newUser = &define.User{}
	var err error
	//首先判断邮箱是否已经被注册过了，如果注册过了，直接返回error
	if dbp.CheckStringExist("User", "Email", Email) {
		err = errors.New("Email has been registered")
		return err
	}
	//如果没有注册，则开始注册
	newUser.UserName = subEmail(Email)
	newUser.Email = Email
	err = dbp.AddNewUser(newUser, Password)
	if err != nil {
		err = errors.New("Sign up failed")
	}
	return err
}

//登录
func (p *Processor)SignIn(signType string, value string, password string) (int, error) {
	if dbp.CheckPassword(signType, value, password) {
		return dbp.GetUserId(signType, value)
	}
	err := errors.New("sign in failed")
	return 0, err
}

//更新用户信息
func (p *Processor)Update(user *define.User, idUser int) error {
	if err := dbp.UpdateUser(idUser, user); err != nil {
		err = errors.New("Update failed")
		return err
	}
	return nil
}

//更新密码
func (p *Processor)UpdatePassword(id int, password string) error {
	if err := dbp.UpdatePassword(id, password); err != nil {
		err = errors.New("Password Update failed")
		return err
	}
	return nil
}

func NewProcessor(dbprocessor *db.Processor) *Processor {
	dbp = dbprocessor
	return &Processor{}
}

func subEmail(email string) string {
	var result string
	for _, r := range email {
		if r == '@' {
			break
		}
		result += string(r)
	}
	return result
}
