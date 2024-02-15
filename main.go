package main

import (
	"fmt"
	"golang.org/x/term"
	"log"
	"net"
	"vimaniac/client"
	clreader "vimaniac/clientreader"
	"vimaniac/flags"
	"vimaniac/message"
	"vimaniac/read"
	"vimaniac/render"
	char "vimaniac/src"
)

// Globals
var w int
var h int

func init() {
	w, h, _ = term.GetSize(0)
	read.Temp = h - 1
	render.MakeRaw()
	render.ClearScreen()
	render.DrawTildas()

	fmt.Print("\x1b[H")
	fmt.Print("\x1b[?25h")
	render.RenderMode(0, h, w)
}
func main() {
	allLines := make([]char.MyCh, 0)
	PtrToChs := new(char.MyCh)
	allLines = append(allLines, *PtrToChs)
	PtrToChs = &allLines[0]
	isClient, ip, filename, err := flags.CheckFlags()
	if err != nil {
		log.Fatal("error(in flags commandline): ", err)
		return
	}
	if len(filename) == 0 {
		filename = append(filename, "")
	}
	a, l := char.Loadfile(filename[0], read.Line, allLines, PtrToChs)
	read.Line = l
	read.Max = l
	render.GetCur(read.Line, PtrToChs.Curx)
	render.RenderLine(read.Line)
	if read.Max < h {
		PtrToChs = &a[read.Line-1]
		render.RenderMode(0, h, w)
		render.GetCur(read.Line, PtrToChs.Curx+8)

	} else if read.Max > h {
		read.Line = 1
		read.Temp = 1
		PtrToChs = &a[0]
		char.RefreshScreen(0, h-1, &a)
		render.RenderMode(0, h, w)
		render.GetCur(read.Temp, PtrToChs.Curx+8)

	}
	if isClient == true {
		recv := client.NewReciver(ip)
		go func() {
			for msg := range recv.Msgch {
				PtrToChs = &a[msg.Line-1]
				message.UnderstandMsg(&a, PtrToChs, msg)
			}
		}()
		conn, err := net.Dial("tcp", recv.DialerAddr)
		defer conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		go recv.ReadLoop(conn)
		clreader.ReadCh(conn, a)
	} else if isClient == false {
		read.ReadChar(a, PtrToChs, false)
	}
	fmt.Print("\x1b[m")
	defer render.Restore()
}
