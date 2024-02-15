package read

import (
	"fmt"
	"golang.org/x/term"
	"log"
	// "net"
	"os"
	// "vimaniac/client"
	"vimaniac/render"
	char "vimaniac/src"
)

var (
	w      int
	h      int
	top    = 0
	bottom int

	Temp int
	Line = 1
	Max  = 1
)

func init() {
	w, h, _ = term.GetSize(0)
	bottom = h - 1
	Temp = h - 1

}

// Globals

func ReadChar(aLns []char.MyCh, ptrch *char.MyCh, isClient bool) {
	// case switchhere
	for {
		// Line1:
		a := make([]byte, 1)
		_, err := os.Stdin.Read(a)
		if err != nil {
			log.Fatal(err)
		}
		switch a[0] {

		case 17:
			// control q
			return
		case 19:
			// save key
			render.GetCur(h, 1)
			filename := render.RenderSave(h, w)
			char.Save(filename, &aLns)
			return
			// maxIsGreater(ptrch)
			// fmt.Print(filename)
		case 127:
			if ptrch.Cols >= 1 && ptrch.Curx >= 1 {
				char.Remove(Line, ptrch)
				maxIsGreater(ptrch)

			}
			// fmt.Print(aLns)
			// control q
		case 27:
			c := make([]byte, 2)
			os.Stdin.Read(c)
			if c[0] == 91 {
				// defer render.RenderMode(0,h, w)
				switch c[1] {
				// tommorrow implement the up and down arrow
				case 65:

					if Line == top+1 && Max > h && top >= 1 && Line > 1 {
						char.ScrollUp(top, bottom, &aLns)
						top--
						bottom--
						Line--
						ptrch = &aLns[Line-1]
					} else if Max > h && Temp >= 1 && Line > 1 {
						Temp--
						Line--
						ptrch = &aLns[Line-1]

					}
					if Line > 1 && Max < h {
						Line--
						ptrch = &aLns[Line-1]

					}
					render.RenderMode(0, h, w)
					maxIsGreater(ptrch)
				case 66:

					if Line == bottom && Max > h && bottom < Max {
						char.ScrollDown(top, bottom, &aLns)
						top++
						bottom++
						Line++
						ptrch = &aLns[Line-1]

					} else if Max > h && Temp < h-1 {
						Temp++
						Line++
						ptrch = &aLns[Line-1]

					} else if Line < Max {
						Line++
						ptrch = &aLns[Line-1]

					}
					render.RenderMode(0, h, w)
					maxIsGreater(ptrch)

				case 68:
					// Temp
					if ptrch.Curx >= 1 {

						ptrch.Curx--
						maxIsGreater(ptrch)
					}
				case 67:

					if ptrch.Curx < ptrch.Cols {
						ptrch.Curx++
						maxIsGreater(ptrch)

					}

				}
			} else {
			outer:
				for {
					render.RenderMode(1, h, w)
					maxIsGreater(ptrch)

					e := make([]byte, 1)
					os.Stdin.Read(e)
					switch e[0] {
					case 'j':
						if Line == bottom && Max > h && bottom < Max {
							char.ScrollDown(top, bottom, &aLns)
							top++
							bottom++
							Line++
							ptrch = &aLns[Line-1]
							// render.RenderMode(0,h, w)
							render.GetCur(Temp, ptrch.Curx+8)

						} else if Max > h && Temp < h-1 {
							Temp++
							Line++
							ptrch = &aLns[Line-1]
							// render.RenderMode(0,h, w)
							render.GetCur(Temp, ptrch.Curx+8)
						} else if Line < Max {
							Line++
							ptrch = &aLns[Line-1]
							// render.RenderMode(0,h, w)
							render.GetCur(Line, ptrch.Curx+8)

						}
					case 'k':
						if Line == top+1 && Max > h && top >= 1 && Line > 1 {
							char.ScrollUp(top, bottom, &aLns)
							top--
							bottom--
							Line--
							ptrch = &aLns[Line-1]
							render.GetCur(Temp, ptrch.Curx+8)
						} else if Max > h && Temp >= 1 && Line > 1 {
							Temp--
							Line--

							ptrch = &aLns[Line-1]
							render.GetCur(Temp, ptrch.Curx+8)

						}
						if Line > 1 && Max < h {
							Line--
							ptrch = &aLns[Line-1]
							render.GetCur(Line, ptrch.Curx+8)

						}
					case 'h':
						if ptrch.Curx >= 1 {

							ptrch.Curx--
							maxIsGreater(ptrch)
						}
					case 'l':

						if ptrch.Curx < ptrch.Cols {
							ptrch.Curx++
							maxIsGreater(ptrch)

						}
					case 'i':
						render.RenderMode(0, h, w)
						maxIsGreater(ptrch)
						break outer

					}
				}

			}
		case '\r':
			Line++
			fmt.Print("\r\n")
			if Line > Max {
				if Line > h-1 {
					char.RenderNext(top, bottom, &aLns)
					render.RenderMode(0, h, w)
					render.GetCur(h-1, ptrch.Curx+8)
					// fmt.Printf("\x1b[%d;%dH", Line, ptrch.Curx+8)

					Temp = h - 1
					top++
					bottom++
					Max = Line

				} else {
					Max = Line

				}

				aLns = append(aLns, char.MyCh{
					Ar:   make([]byte, 0),
					Cols: 0,
					Curx: 0,
				})
				ptrch = &aLns[Line-1]
				// char.NewLine(&aLns, Line, ptrch)
				// }
				render.RenderLine(Line)
			}
			if Line < Max {

				char.InsertLine(Line, &aLns)
				ptrch = &aLns[Line-1]
				Max++
				if Max < h {

					char.RefreshScreen(0, Max, &aLns)
					render.RenderMode(0, h, w)
					render.GetCur(Line, ptrch.Curx+8)

				} else if Max >= h {
					Temp++
					char.RefreshScreen(top, bottom, &aLns)
					render.RenderMode(0, h, w)
					render.GetCur(Temp, ptrch.Curx+8)
				}
			}

			// aLns[Line] := MyCh
			// fmt.Print(bottom)
		default:
			// insert directly to allLines
			character := char.Insert(Line, a, ptrch)
			fmt.Print(character)
			maxIsGreater(ptrch)
			// fmt.Print(aLns)
			// ptrch.Curx++
			// fmt.Print(aLns)

		}
		if a[0] == 17 {
			return
		}

	}

}

// refactor by adding func isMaxH takes Max produces bool

func maxIsGreater(ch *char.MyCh) {
	if Max < h {

		render.GetCur(Line, ch.Curx+8)
	} else if Max > h {
		// fmt.Print("Temp")
		render.GetCur(Temp, ch.Curx+8)

	}

}
