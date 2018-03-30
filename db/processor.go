package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"goservertest/define"
	"time"
	"goservertest/util"
	"encoding/json"
)

var DB *sql.DB

type Processor struct {
}

//新增用户
func (p *Processor) AddNewUser(user *define.User, password string) error {
	var err error = nil

	stmt, err := DB.Prepare(`INSERT user (Username, Email, Password, CreateTime) values (?,?,?,?)`)
	if err != nil {
		log.Print(err)
		return err
	}
	res, err := stmt.Exec(user.UserName, user.Email, password, time.Now())
	if err != nil {
		log.Print(err)
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		log.Print(err)
		return err
	}
	return err
}

//更新用户信息
func (p *Processor) UpdateUser(UserId int, user *define.User) error {
	var err error = nil
	stmt, err := DB.Prepare(`UPDATE User SET UserName=?,Telephone=?,Email=?,Sex=? WHERE idUser=?`)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(user.UserName, user.Telephone, user.Email, user.Sex, UserId)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err
}

//修改密码
func (p *Processor) UpdatePassword(UserId int, password string) error {
	var err error = nil
	stmt, err := DB.Prepare(`UPDATE User SET Password=? WHERE idUser =?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(password, UserId)
	return err
}
func (p *Processor) TesConn() {
	log.Println("connect successfully")
}

//Game
func (p *Processor) CreateNewGame(game *define.Game) (int, error) {
	var err error = nil
	stmt, err := DB.Prepare(`INSERT Game (StartTime, Endtime, BlackUser, WhiteUser, BoardSize, Board, GameRecord) VALUES (?,?,?,?,?,?,?)`)
	if err != nil {
		log.Println("1--", err)
		return -1, err
	}

	res, err := stmt.Exec(
		time.Now(),
		time.Now(),
		game.BlackUser,
		game.WhiteUser,
		game.BoardSize,
		game.Board,
		game.GameLog.Encode())
	if err != nil {
		log.Println("2--", err)
		return -1, err
	}

	num, err := res.LastInsertId()
	if err != nil {
		log.Println("3--", err)
		return -1, err
	}
	return int(num), err
}

//从数据库读取棋局信息
func (p *Processor) GetGame(id int) *define.Game {
	var game = &define.Game{}
	game.GameLog = util.NewGameLog()
	var err error

	stmt, err := DB.Prepare(`SELECT
	StartTime,
	EndTime,
	IsEnd,
	BlackUser,
	WhiteUser,
	Winner,
	BlackScore,
	WhiteScore,
	BoardSize,
	Board,
	TotalPoint,
	GameRecord
	FROM Game WHERE idGame =?`)
	if err != nil {
		log.Println("1--", err)
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println("2--", err)
	}
	for rows.Next() {
		var tempTime1, tempTime2 string
		var templog []uint8
		err = rows.Scan(
			&tempTime1,
			&tempTime2,
			&game.IsEnd,
			&game.BlackUser,
			&game.WhiteUser,
			&game.Winner,
			&game.BlackScore,
			&game.WhiteScore,
			&game.BoardSize,
			&game.Board,
			&game.TotalPoint,
			&templog)
		if err != nil {
			log.Println("3--", err)
		}
		game.StartTime, _ = time.Parse("2006-01-02 15:04:05", tempTime1)
		game.EndTime, _ = time.Parse("2006-01-02 15:04:05", tempTime2)
		json.Unmarshal(templog, game.GameLog)
	}
	return game
}

//更新棋局信息
func (p *Processor) UpdateGame(id int, game *define.Game) {
	endTime := game.EndTime
	totalpoint := game.TotalPoint
	board := game.Board
	glog := game.GameLog

	stmt, err := DB.Prepare(`UPDATE game SET Endtime=?, TotalPoint=?, Board=?, GameRecord=? WHERE idGame=?`)
	if err != nil {
		log.Println("1--",err)
	}
	res, err := stmt.Exec(endTime, totalpoint, board, glog.Encode(), id)
	if err != nil {
		log.Println("2--",err)
	}
	_, err = res.RowsAffected()
	if err != nil {
		log.Println("3--",err)
	}
	return
}

//创建数据库处理器，暂时不知道需要啥参数
func NewProcessor() *Processor {
	var err error

	if DB, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/goDBTest?charset-utf8"); err != nil {
		log.Println("Open log", err)
		return nil
	}

	if err = DB.Ping(); err != nil {
		log.Println("Ping log", err)
		return nil
	}

	return &Processor{
	}
}

//检查某字符串字段是否存在
func (p *Processor) CheckStringExist(table string, checkItem string, value string) bool {
	var err error
	var msg string
	quary := "SELECT 1 FROM " + table + " WHERE " + checkItem + " = \"" + value + "\""
	if err = DB.QueryRow(quary).Scan(&msg); err != nil {
		log.Println("check log ", err)
		return false
	}
	return true
}

//检查某数值字段是否存在
func (p *Processor) CheckIntExist(table string, checkItem string, value string) bool {
	var err error
	var msg string
	quary := "SELECT 1 FROM " + table + " WHERE " + checkItem + " = " + value
	if err = DB.QueryRow(quary).Scan(&msg); err != nil {
		log.Println("check log ", err)
		return false
	}
	return true
}

//检查密码是否正确
func (p *Processor) CheckPassword(signType string, value string, password string) bool {
	var oldPassword string
	var err error
	stmt, err := DB.Prepare("SELECT Password FROM User WHERE " + signType + " =?")
	if err != nil {
		log.Print(err)
		return false
	}
	if err = stmt.QueryRow(value).Scan(&oldPassword); err != nil {
		log.Println(err)
		return false
	}
	return password == oldPassword
}

//得到用户id
func (p *Processor) GetUserId(signType string, value string) (int, error) {
	var id int
	var err error
	stmt, err := DB.Prepare("SELECT idUser FROM User WHERE " + signType + " =?")
	if err != nil {
		log.Print(err)
		return 0, err
	}
	if err = stmt.QueryRow(value).Scan(&id); err != nil {
		log.Println(err)
		return 0, err
	}
	return id, nil
}
