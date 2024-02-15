package render

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"golang.org/x/term"
)

var Esp string = "\x1b"

func MakeRaw() {
	_, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatal(err)
	}

}

func RenderLine(line int) {
	fmt.Print(Esp + "[0m")
	fmt.Print(Esp + "[1000D")
	fmt.Print(Esp + "[48;5;194m") //color backgroud
	fmt.Print(Esp + "[2m")        // Bold
	fmt.Print(Esp + "[38;5;94m")
	fmt.Print("|")
	length := len(strconv.Itoa(line))
	spaces := 4 - length
	fmt.Printf("%*s", spaces, "")

	fmt.Print(strconv.Itoa(line))

	fmt.Print(" ")
	fmt.Print(Esp + "[0m")
	fmt.Print(" ")
	CustomizeTerm()

}
func RenderSave(h int, w int) string {
	fmt.Print(Esp + "[0m")
	fmt.Print(Esp + "[1000D")
	fmt.Print(Esp + "[48;5;216m") //color backgroud
	fmt.Print(Esp + "[2m")        // Bold
	fmt.Print(Esp + "[38;5;94m")
	file := "Enter Filename:"

	fmt.Print("\x1b[5 q")
	GetCur(h, 1)
	fmt.Printf("%*s", w, "")
	GetCur(h, 1)
	fmt.Print(file)
	GetCur(h, len(file)+2)
	var input string
	for {

		a := make([]byte, 1)
		_, err := os.Stdin.Read(a)
		if err != nil {
			log.Fatal(err)
		}
		if a[0] == '\r' || a[0] == '\n' {
			break

		}
		if a[0] == 27 {
			continue
		}
		if a[0] == 127 && len(input) >= 1 {

			fmt.Print("\033[1D \033[1D")
			input = input[:len(input)-1]
			if len(input) == 0 {
				fmt.Print(" ")

			}
		} else {

			input += string(a[:])
		}

		fmt.Print(string(a[:]))

	}
	if len(input) == 0 {
		input = "VimIsGreat"
	}
	fmt.Print(Esp + "[0m")
	CustomizeTerm()
	return input
}
func RenderMode(Mode int, h int, w int) {
	fmt.Print(Esp + "[0m")
	fmt.Print(Esp + "[1000D")
	fmt.Print(Esp + "[48;5;216m") //color backgroud
	fmt.Print(Esp + "[2m")        // Bold
	fmt.Print(Esp + "[38;5;94m")
	insert := "~~~~~~~~~~INSERT~~~~~~~~~~"
	command := "~~~~~~~~~~COMMAND~~~~~~~~~~"
	// var i string
	switch Mode {
	case 0:
		fmt.Print("\x1b[5 q")
		GetCur(h, 1)
		fmt.Printf("%*s", (w-len(insert))/2, "")
		fmt.Print(insert)
		fmt.Printf("%*s", (w-len(insert))/2, "")
	case 1:
		fmt.Print("\x1b[1 q")
		GetCur(h, 1)
		fmt.Printf("%*s", (w-len(command))/2, "")
		fmt.Print(command)
		fmt.Printf("%*s", (w-len(command))/2, "")
	}
	fmt.Print(Esp + "[0m")
	CustomizeTerm()

}
func GetCur(line, crx int) {

	fmt.Printf("%s[%d;%dH", Esp, line, crx)
}

func DrawTildas() {
	_, h, e := term.GetSize(0)
	if e != nil {
		fmt.Print("Error", e)
	}
	fmt.Print("\r\n")

	for i := 0; i < h-1; i++ {
		fmt.Print("~\r\n")

	}

}

func SetCurColor() {

}
func RenderLoading() {

	fmt.Print(Esp + "[?25l")
	for i := 0; i < 3; i++ {
		fmt.Print(".")
		// go time.Sleep(1 * time.Second)

	}

	fmt.Print("\033[1D \033[1D")
	fmt.Print("\033[1D \033[1D")
	fmt.Print("\033[1D \033[1D")
	// fmt.Print(Esp + "[?25h")

}
func ClearScreen() {
	fmt.Print(Esp + "[2J")
	fmt.Print(Esp + "[H")
	fmt.Print(Esp + "[?25l")

}
func Restore() {
	fmt.Print("\x1b[2J")
	fmt.Print("\x1b[H")
	fmt.Print("\x1b[m")
	fmt.Print("\x1b[1 q")
	fmt.Print(Esp + "[?25h")
	exec.Command("reset").Run()

}
func CustomizeTerm() {
	fmt.Print(Esp + "[38;5;3m")

	fmt.Print(Esp + "[10m")
	fmt.Print(Esp + "[3m")

}
