package main

import (
	"goservertest/db"
	"goservertest/define"
	"time"
	"goservertest/util"
)

func main() {
	dbp := db.NewProcessor()
	game := &define.Game{}
	id := 1

	n := make([][]uint8, 9*9)
	for k := range n {
		n[k] = make([]uint8, 9)
	}
	for i := 0; i < len(n); i++ {
		n[i/9][i%9] = uint8(i)
	}

	game.EndTime = time.Now()
	game.TotalPoint++
	game.Board = util.Step2Log(n, 9)
	game.GameLog = util.NewGameLog()
	game.GameLog.AddSingleLog(util.NewSinglelog(1, 1, game.Board, nil))

	dbp.UpdateGame(id, game)
}
