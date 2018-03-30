package main

import (
	"fmt"
	"goservertest/db"
	//"goservertest/game"
)

func main() {
	dbp := db.NewProcessor()
	//game := game.NewProcessor(dbp)
	//requsetId := 13
	//BoardSize := 9
	//num, err := game.StartNewGame(requsetId, BoardSize)
	//if err != nil {
	//	fmt.Println("error")
	//} else {
	//	fmt.Println(num)
	//}

	game1 := dbp.GetGame(3)
	fmt.Println(game1.BoardSize)
}
