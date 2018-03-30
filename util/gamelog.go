package util

import (
	"encoding/json"
	"log"
)

type SingleLog struct {
	Player uint8
	Place  uint16
	Board  []uint8
	Lose   []uint16
}

type GameLog struct {
	Logs []SingleLog
}

func (gl *GameLog) Encode() []byte {
	r, err := json.Marshal(gl)
	if err != nil {
		log.Print(err)
		return nil
	}
	return r
}

func (gl *GameLog) Decode(source []byte) {
	json.Unmarshal(source, gl)
}

func (gl *GameLog) AddSingleLog(sl *SingleLog) {
	gl.Logs = append(gl.Logs, *sl)
}

func NewGameLog() *GameLog {
	gl := &GameLog{}
	gl.Logs = make([]SingleLog, 0)
	return gl
}

func NewSinglelog(Player uint8, Place uint16, Board []uint8, Lose []uint16) *SingleLog {
	sl := &SingleLog{Player,Place,Board,Lose}
	return sl
}
