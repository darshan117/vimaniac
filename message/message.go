package message

import (
	"fmt"
	"golang.org/x/term"
	"vimaniac/client"
	"vimaniac/render"
	char "vimaniac/src"
)

var (
	w      int
	h      int
	top    = 0
	bottom int
	Temp   int
	Line   = 1
	Max    = 1
)

func init() {
	w, h, _ = term.GetSize(0)
	bottom = h - 1
	Temp = h - 1

}

// Globals

func UnderstandMsg(aLns *[]char.MyCh, ptrch *char.MyCh, mg client.Message) {
	switch mg.Msg {

	case "exit":
		return
	case "remove":
		if ptrch.Cols >= 1 && ptrch.Curx >= 1 {
			char.Remove(Line, ptrch)
			maxIsGreater(ptrch)

		}
	case "up":

		if Line == top+1 && Max > h && top >= 1 && Line > 1 {
			char.ScrollUp(top, bottom, &(*aLns))
			top--
			bottom--
			Line--
			ptrch = &(*aLns)[Line-1]
		} else if Max > h && Temp >= 1 && Line > 1 {
			Temp--
			Line--
			ptrch = &(*aLns)[Line-1]

		}
		if Line > 1 && Max < h {
			Line--
			ptrch = &(*aLns)[Line-1]

		}
		render.RenderMode(0, h, w)
		maxIsGreater(ptrch)
	case "down":

		if Line == bottom && Max > h && bottom < Max {
			char.ScrollDown(top, bottom, &(*aLns))
			top++
			bottom++
			Line++
			ptrch = &(*aLns)[Line-1]

		} else if Max > h && Temp < h-1 {
			Temp++
			Line++
			ptrch = &(*aLns)[Line-1]

		} else if Line < Max {
			Line++
			ptrch = &(*aLns)[Line-1]

		}
		render.RenderMode(0, h, w)
		maxIsGreater(ptrch)

	case "left":
		if ptrch.Curx >= 1 {

			ptrch.Curx--
			maxIsGreater(ptrch)
		}
	case "right":

		if ptrch.Curx < ptrch.Cols {
			ptrch.Curx++
			maxIsGreater(ptrch)

		}

	case "newline":
		fmt.Print("\r\n")
		Line++
		if Line > Max {
			if Line > h-1 {
				char.RenderNext(top, bottom, &(*aLns))
				render.RenderMode(0, h, w)
				render.GetCur(h-1, ptrch.Curx+8)
				Temp = h - 1
				top++
				bottom++
				Max = Line
			} else {
				Max = Line
			}
			(*aLns) = append((*aLns), char.MyCh{
				Ar:   make([]byte, 0),
				Cols: 0,
				Curx: 0,
			})
			render.RenderLine(Line)
		}
		if Line < Max {
			char.InsertLine(Line, &(*aLns))
			Max++
			if Max < h {
				char.RefreshScreen(0, Max, &(*aLns))
				render.RenderMode(0, h, w)
				render.GetCur(Line, ptrch.Curx+8)

			} else if Max >= h {
				Temp++
				char.RefreshScreen(top, bottom, &(*aLns))
				render.RenderMode(0, h, w)
				render.GetCur(Temp, ptrch.Curx+8)
			}
		}
	case "insert":
		character := char.Insert(Line, mg.Ch, ptrch)
		fmt.Print(character)
		maxIsGreater(ptrch)
	}

}
func maxIsGreater(ch *char.MyCh) {
	if Max < h {
		render.GetCur(Line, ch.Curx+8)
	} else if Max > h {
		render.GetCur(Temp, ch.Curx+8)
	}
}
