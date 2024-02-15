package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Receiver struct {
	DialerAddr string
	Quitch     chan struct{}
	Msgch      chan Message
}
type Message struct {
	Msg  string `json:"operation"`
	Line int    `json:"line"`
	Ch   []byte `json:"character"`
	Cury int    `json:"cury"`
}

func NewReciver(dialerAddr string) *Receiver {
	return &Receiver{
		DialerAddr: dialerAddr,
		Quitch: make(chan struct {
		}),
		Msgch: make(chan Message),
	}

}
func (s *Receiver) ReadLoop(r net.Conn) {
	buf := make([]byte, 2048)
	for {
		n, err := r.Read(buf)
		if err != nil {
			log.Println(err)
			r.Close()
		}
		msg := buf[:n]
		var data Message
		if err := json.Unmarshal(msg, &data); err != nil {
			log.Fatal("unmarshal error ", err)
		}
		s.Msgch <- data
	}
}
func (s *Receiver) Dialer() {
	conn, err := net.Dial("tcp", s.DialerAddr)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("conneted succesfully to ", s.DialerAddr)
	go s.ReadLoop(conn)
}
func ModMessage(conn net.Conn, operation string, line int, ch []byte) {
	mesg := Message{Msg: operation, Line: line, Ch: ch}
	jsonMsg, err := json.Marshal(mesg)
	if err != nil {
		log.Fatal("error in json : ", err)
	}
	conn.Write(jsonMsg)
}
