package util

import "bytes"

func Step2Log(source [][]byte, size int) []byte {
	temp := bytes.Join(source, []byte(""))
	return temp[: size*size]
}

func Log2Step(source []byte, size int) [][]byte {
	ans := make([][]byte, size*size)
	for k := range ans {
		ans[k] = make([]byte, size)
		copy(ans[k], source[size*k:size*(k+1)])
	}
	return ans
}

func AddPiece(source [][]byte, place uint16, player uint8) [][]byte {
	x := place / 19
	y := place % 19
	source[x][y] = player
	return source
}
