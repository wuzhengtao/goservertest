package game

import (
	"goservertest/db"
	"math/rand"
	"time"
	"goservertest/define"
	"goservertest/util"
	"goservertest/rule"
)

var dbp *db.Processor

type Processor struct {
}

//申请开局
/**
首先，我要申请者的id，然后要棋盘大小
返回的是一个已经进入数据库的棋局id

过程
	随机决定执方
	创建棋盘
	创建日志文件
	存入数据库
 */
func (p *Processor) StartNewGame(requestId int, BoardSize int) (int, error) {
	var newgame = &define.Game{}

	player := RandomPlayer()
	if player == 0 {
		newgame.BlackUser = requestId
		newgame.WhiteUser = -1
	} else {
		newgame.WhiteUser = requestId
		newgame.BlackUser = -1
	}

	newgame.Board = make([]uint8, BoardSize*BoardSize)
	for k := range newgame.Board {
		newgame.Board[k] = uint8(k)
	}

	newgame.BoardSize = BoardSize

	newgame.GameLog = util.NewGameLog()

	return dbp.CreateNewGame(newgame)

}

/**
下棋步骤
1-申请对局之后先读取上一棋局
2-落子合理性的审核暂定放在客户端，默认服务端的落子全部正确，此步跳过
3-调用围棋规则，获取最新棋盘以及被提子的记录
4-将此步棋的log更新
5-更新数据库
 */
func (p *Processor) Play(gameid int, player uint8, place uint16) [][]uint8 {
	//根据id读取棋盘
	game := dbp.GetGame(gameid)
	size := game.BoardSize

	//准备下棋
	board := util.Log2Step(game.Board, size)
	board = util.AddPiece(board, place, player)
	newboard, _ := rule.GameCenterLogic(board, player, uint8(game.BoardSize))

	newboardLog := util.Step2Log(newboard, size)

	//更新log
	sl := util.NewSinglelog(player, place, newboardLog, nil)
	game.GameLog.AddSingleLog(sl)

	//更新部分数据
	game.EndTime = time.Now()
	game.TotalPoint++
	game.Board = newboardLog

	dbp.UpdateGame(gameid, game)
	return newboard
}

func NewProcessor(dbprocessor *db.Processor) *Processor {
	dbp = dbprocessor
	return &Processor{}
}

/**
随机数
返回0则是黑方
1则是白方
 */
func RandomPlayer() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ran := r.Intn(2)
	return ran
}
