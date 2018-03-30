package gateway

import (
	"goservertest/db"
	"goservertest/login"
	"goservertest/game"
	"net"
	"log"
	"goservertest/define"
	"time"
	"fmt"
	"encoding/json"
)

var dbp *db.Processor
var lp *login.Processor
var gp *game.Processor

type Processor struct {
}

//这个函数负责启动整个服务器，并且负责所有的数据收发
func (p *Processor) Start() {
	listen_sock, err := net.Listen(define.TCP, define.HOST)
	checkError(err)
	defer listen_sock.Close()

	for {
		new_conn, err := listen_sock.Accept()
		c := make(chan []uint8, 10)
		if err != nil {
			continue
		}

		go recvConnMsg(new_conn, c)
	}
}

//单纯收发信息，不负责信息的解析
func recvConnMsg(conn net.Conn, c chan []uint8) {
	//设置当客户端3分钟内无数据请求时，自动关闭conn
	conn.SetReadDeadline(time.Now().Add(time.Minute * 3))
	defer conn.Close()
	buf := make([]byte, 1000)

	lproxy := login.NewProxy(lp, c)

	defer conn.Close()

	go sendMessage(conn, c)

	for {
		n, err := conn.Read(buf)

		if err != nil {
			continue
		}
		rec := buf[0:n]

		message := &define.Message{}
		json.Unmarshal(rec, message)

		code := message.Code
		msg := message.Msg
		switch code / 100 {
		case 1:
			lproxy.ProcessMessage(code, msg)
		default:

		}
	}

	fmt.Println("conn ended")
}

func sendMessage(conn net.Conn, ch chan []uint8) {
	for {
		msg := <-ch
		conn.Write(msg)
	}
}

func NewProcessor(dbprocessor *db.Processor, lprocessor *login.Processor, gprocessor *game.Processor) *Processor {
	dbp = dbprocessor
	lp = lprocessor
	gp = gprocessor
	return &Processor{}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
