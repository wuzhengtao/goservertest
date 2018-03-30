package rule

import (
	"goservertest/define"
)

// 围棋提子逻辑

func GameCenterLogic(gameInfo [][]uint8, getColor uint8, size uint8) ([][]uint8, []int) {
	getColor = 3 - getColor
	ko := make([]int, 2)
	TakeDeadChess(gameInfo, getColor, size, ko)
	return gameInfo, ko
}

func TakeDeadChess(gameInfo [][]uint8, getColor uint8, gameSize uint8, ko []int) {
	count := 0
	allInfo := make([][2]uint8, gameSize*gameSize)
	hasDealWith := make([][]bool, gameSize)
	for k := range hasDealWith {
		hasDealWith[k] = make([]bool, gameSize)
	}
	var (
		top int
		pos int
	)
	for i := uint8(0); i < (gameSize); i++ {
		for j := uint8(0); j < (gameSize); j++ {
			if gameInfo[i][j] != getColor || hasDealWith[i][j] {
				continue
			}
			hasDealWith[i][j] = true
			allInfo[top][0] = i
			allInfo[top][1] = j
			top++
			for pos < top {
				x := allInfo[pos][0]
				y := allInfo[pos][1]
				if x > 0 && !hasDealWith[x-1][y] && getColor == gameInfo[x-1][y] {
					hasDealWith[x-1][y] = true
					allInfo[top][0] = x - 1
					allInfo[top][1] = y
					top++
				}
				if x < (gameSize-1) && !hasDealWith[x+1][y] && getColor == gameInfo[x+1][y] {
					hasDealWith[x+1][y] = true
					allInfo[top][0] = x + 1
					allInfo[top][1] = y
					top++
				}
				if y > 0 && !hasDealWith[x][y-1] && getColor == gameInfo[x][y-1] {
					hasDealWith[x][y-1] = true
					allInfo[top][0] = x
					allInfo[top][1] = y - 1
					top++
				}
				if y < (gameSize-1) && !hasDealWith[x][y+1] && getColor == gameInfo[x][y+1] {
					hasDealWith[x][y+1] = true
					allInfo[top][0] = x
					allInfo[top][1] = y + 1
					top++
				}
				pos++ //next step
			}
			if top > 0 && IsDeadChess(gameInfo, allInfo, top, gameSize) {
				ClearDeadChess(gameInfo, allInfo, top, ko)
				count += top
			}
			top = 0
			pos = 0
		}
	}
	if count != 1 {
		ko[0] = -1
		ko[1] = -1
	}
}

func IsDeadChess(gameInfo [][]uint8, allInfo [][2]uint8, top int, gameSize uint8) bool {
	var (
		x, y uint8
	)
	for i := 0; i < top; i++ {
		x = allInfo[i][0]
		y = allInfo[i][1]
		if x > 0 && gameInfo[x-1][y] == define.SPACE_G {
			return false
		}
		if x < gameSize-1 && gameInfo[x+1][y] == define.SPACE_G {
			return false
		}
		if y > 0 && gameInfo[x][y-1] == define.SPACE_G {
			return false
		}
		if y < gameSize-1 && gameInfo[x][y+1] == define.SPACE_G {
			return false
		}
	}
	return true
}

// 清除死子
func ClearDeadChess(gameInfo [][]uint8, allInfo [][2]uint8, top int, ko []int) {
	for i := 0; i < top; i++ {
		gameInfo[allInfo[i][0]][allInfo[i][1]] = define.SPACE_G
		ko[0] = int(allInfo[i][0])
		ko[1] = int(allInfo[i][1])
	}
}
