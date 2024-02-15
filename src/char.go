package char

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"vimaniac/render"
)

// https://darrenburns.net/posts/piece-table/
type MyCh struct {
	Ar   []byte
	Cols int
	Curx int
}

func Insert(Line int, char []byte, ch *MyCh) string {
	lastelm := len(ch.Ar)
	v := ""
	if ch.Cols == ch.Curx {
		ch.Ar = append(ch.Ar, char[0])
		ch.Cols++
		ch.Curx++
		return string(ch.Ar[lastelm])
	}
	if ch.Curx < ch.Cols {
		index := ch.Curx
		ch.Ar = append(ch.Ar[:index+1], ch.Ar[index:]...)
		ch.Ar[index] = char[0]
		for _, letters := range ch.Ar[index:] {
			v += string(letters)
		}
		ch.Cols++
		ch.Curx++
		return v

	}
	return ""

}
func Remove(Line int, ch *MyCh) {
	lastelm := len(ch.Ar) - 1
	if ch.Cols == ch.Curx {
		ch.Ar = ch.Ar[:lastelm]
		fmt.Print("\033[1D \033[1D")
		ch.Cols--
		ch.Curx--
	}
	if ch.Curx < ch.Cols {
		index := ch.Curx
		ch.Ar = append(ch.Ar[:index-1], ch.Ar[index:]...)
		fmt.Print("\033[1D \033[1D")
		for _, letters := range ch.Ar[index-1:] {
			fmt.Print(string(letters))
		}
		fmt.Print(" ")
		ch.Cols--
		ch.Curx--
	}
}

func InsertLine(Line int, aa *[]MyCh) {
	(*aa) = append((*aa)[:Line], (*aa)[Line-1:]...)
	(*aa)[Line-1] = MyCh{
		Ar:   make([]byte, 0),
		Cols: 0,
		Curx: 0,
	}
}
func RefreshScreen(t int, b int, ar *[]MyCh) {
	render.ClearScreen()
	defer fmt.Print("\x1b[?25h")
	n := t
	for _, l := range (*ar)[t:b] {
		render.RenderLine(n + 1)
		for _, c := range l.Ar {
			fmt.Print(string(c))

		}
		fmt.Print("\r\n")
		n++
	}

}

// if top Line
func ScrollDown(t int, b int, ar *[]MyCh) {
	render.ClearScreen()
	defer fmt.Print("\x1b[?25h")
	n := t + 1
	for _, l := range (*ar)[t+1 : b+1] {
		render.RenderLine(n + 1)
		for _, c := range l.Ar {
			fmt.Print(string(c))

		}
		fmt.Print("\r\n")
		n++
	}

}
func ScrollUp(t int, b int, ar *[]MyCh) {
	render.ClearScreen()
	defer fmt.Print("\x1b[?25h")
	n := t - 1
	for _, l := range (*ar)[t-1 : b-1] {
		render.RenderLine(n + 1)
		for _, c := range l.Ar {
			fmt.Print(string(c))
		}
		fmt.Print("\r\n")
		n++
	}

}
func RenderNext(t int, b int, ar *[]MyCh) {
	render.ClearScreen()
	defer fmt.Print("\x1b[?25h")
	n := t + 1
	for _, l := range (*ar)[t+1 : b] {
		render.RenderLine(n + 1)
		for _, c := range l.Ar {
			fmt.Print(string(c))
		}
		fmt.Print("\r\n")
		n++
	}
}
func NewLine(aLns *[]MyCh, Line int, ptrch *MyCh) {
	*aLns = append(*aLns, MyCh{
		Ar:   make([]byte, 0),
		Cols: 0,
		Curx: 0,
	})
	render.RenderLine(Line)
	fmt.Print(reflect.TypeOf(*aLns))
}
func Save(name string, a *[]MyCh) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, v := range *a {
		_, err := f.Write(v.Ar)

		if err != nil {
			log.Fatal(err)
		}
		f.WriteString("\n")

	}

}
func Loadfile(filename string, Line int, a []MyCh, ch *MyCh) ([]MyCh, int) {
	if len(flag.Args()) > 0 {
		_, err := os.Stat(filename)
		if err != nil {
			log.Fatal("file does not exist")
		}
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)

		}
		fmt.Print("Loading")
		for {
			b := make([]byte, 1)
			render.RenderLoading()
			n, err := f.Read(b)
			if err != nil {
				if err == io.EOF {
					fmt.Print("\x1b[?25h")
					fmt.Print("\x1b[0m")
					render.CustomizeTerm()
					return a, Line

				}
			}
			if n < len(b) {
				break
			}
			if b[0] == 13 || b[0] == 10 {
				Line++
				a = append(a, MyCh{
					Ar:   make([]byte, 0),
					Cols: 0,
					Curx: 0,
				})
				ch = &a[Line-1]
			} else if b[0] != '\n' {
				Insert(Line, b, ch)
			} else if b[0] == 9 {
				b[0] = 13
				for i := 0; i < 4; i++ {
					Insert(Line, b, ch)
				}
			}
		}
	}
	return a, Line
}
