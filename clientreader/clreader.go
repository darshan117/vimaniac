package clreader

import (
	"log"
	"net"
	"os"
	"vimaniac/client"
	"vimaniac/message"
	char "vimaniac/src"
)

func ReadCh(conn net.Conn, aLns []char.MyCh) {
	for {
		a := make([]byte, 1)
		_, err := os.Stdin.Read(a)
		if err != nil {
			log.Fatal(err)
		}
		switch a[0] {
		case 17:
			return
		case 127:
			client.ModMessage(conn, "remove", message.Line, a)
		case 27:
			c := make([]byte, 2)
			os.Stdin.Read(c)
			if c[0] == 91 {
				switch c[1] {
				case 65:
					client.ModMessage(conn, "up", message.Line, a)
				case 66:
					client.ModMessage(conn, "down", message.Line, a)
				case 68:
					client.ModMessage(conn, "left", message.Line, a)
				case 67:
					client.ModMessage(conn, "right", message.Line, a)
				}
			} else {
			outer:
				for {
					e := make([]byte, 1)
					os.Stdin.Read(e)
					switch e[0] {
					case 'j':
						client.ModMessage(conn, "down", message.Line, a)
					case 'k':
						client.ModMessage(conn, "down", message.Line, a)
					case 'h':
						client.ModMessage(conn, "left", message.Line, a)
					case 'l':
						client.ModMessage(conn, "right", message.Line, a)
					case 'i':
						break outer
					}
				}
			}
		case '\r':
			client.ModMessage(conn, "newline", message.Line, a)
		default:
			client.ModMessage(conn, "insert", message.Line, a)
		}
		if a[0] == 17 {
			client.ModMessage(conn, "exit", message.Line, a)
		}
	}
}
