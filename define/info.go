package define

import (
	"image"
	"time"
	"goservertest/util"
)

type User struct {
	UserName           string      //用户名
	UserIcon           image.Image //头像
	Telephone          string      //电话
	Email              string      //邮箱
	Sex                int         //性别
	TotalGameNormal    int         //大棋盘总局数
	TotalVectoryNormal int         //大棋盘胜局数
	TotalGameMidium    int         //中棋盘总局数
	TotalVectoryMidium int         //中棋盘胜局数
	TotalGameSmall     int         //小棋盘总局数
	TotalVectorySmall  int         //小棋盘胜局数
}

type Game struct {
	StartTime  time.Time     //棋局开始时间
	EndTime    time.Time     //棋局结束时间
	IsEnd      uint8         //判断棋局是否结束
	BlackUser  int           //黑方
	WhiteUser  int           //白方
	Winner     int           //赢方
	BlackScore int           //黑方分数
	WhiteScore int           //白方分数
	Board      []uint8       //目前棋局形势
	BoardSize  int           //棋盘大小
	TotalPoint int           //总目数
	GameLog    *util.GameLog //棋局记录
}

type Message struct {
	Code int     //请求代码
	Msg  []uint8 //请求信息
}

type ErrorMsg struct {
	Code int    //错误代码
	Err  string //错误信息
}
