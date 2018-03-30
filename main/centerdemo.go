package main

import (
	"fmt"
	"goservertest/rule"
)

func main() {
	size := uint8(9)
	gameInfo := make([][]uint8, size)
	ko := make([]int, 2)
	for k := range gameInfo {
		gameInfo[k] = make([]uint8, size)
	}
	PrintInfo(gameInfo)
	var (
		x int
		y int
		player = uint8(1)
	)
	for true {
		fmt.Println("请输入 x 和 y 值，0～8")
		fmt.Scanf("%d%d", &x, &y)
		if x == ko[0] && y == ko[1] {
			fmt.Println("此为禁着点")
			continue
		}
		gameInfo[x][y] = player
		gameInfo, ko = rule.GameCenterLogic(gameInfo, player, size)
		PrintInfo(gameInfo)
		player = 3 - player
	}
}

func PrintInfo(game [][]uint8) {
	for i := range game{
		for j := range game[i] {
			fmt.Print(game[i][j], " ")
		}
		fmt.Println()
	}
}