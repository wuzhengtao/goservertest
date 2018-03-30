package main

import (
	"fmt"
	"goservertest/util"
)

func main() {
	n := make([][]uint8, 9*9)
	for k := range n {
		n[k] = make([]uint8, 9)
	}
	for i := 0; i < len(n); i++ {
		n[i/9][i%9] = uint8(i)
	}
	fmt.Println(n)
	fmt.Println(len(n), cap(n))

	n1 := util.Step2Log(n, 9)
	fmt.Println(n1)
	fmt.Println(len(n1), cap(n1))

	n2:= util.Log2Step(n1, 9)
	fmt.Println(n2)
	fmt.Println(len(n2), cap(n2))
}
