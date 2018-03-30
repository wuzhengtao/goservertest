package login

/**
json 负责解析与打包对应的信息

 */

type SignUpMsg struct {
	Email    string
	Password string
}

type SignInMsg struct {
	SignType string
	Value    string
	Password string
}

type UpdateMsg struct {
	Id        int
	UserName  string
	Telephone string
	Email     string
	Sex       int
}

type UpdatePW struct {
	Id        int
	Password  string
}
