package define

//性别
const (
	MALE   = 1 //男性
	FOMALE = 2 //女性
	SECRET = 3 //保密
)

//棋盘大小
const (
	NORMAL = 19 //大
	MIDIUM = 13 //中
	SMALL  = 9  //小
)

//胜负
const (
	BLACKWIN = 1 //黑方赢
	WHITEWIN = 2 //白方赢
	DRAW     = 3 //和棋
)

//登录类型
const (
	SIGN_USER_NAME = "UserName"  //用户名
	SIGN_TELEPHONE = "Telephone" //手机号
	SIGN_EMAIL     = "Email"     //邮箱
)

//棋盘信息
const (
	SPACE_G = 0 //二进制为00
	BLACK_G = 2 //二进制为10
	WHITE_G = 3 //二进制为11
)

//网络相关
const (
	TCP  = "tcp"
	HOST = "localhost:10000"
)
