package main

import (
	db2 "goservertest/db"
	"goservertest/login"
	"goservertest/game"
	"goservertest/gateway"
)

func main()  {
	db := db2.NewProcessor()
	lp := login.NewProcessor(db)
	gp := game.NewProcessor(db)

	gateway := gateway.NewProcessor(db, lp, gp)
	gateway.Start()
}
