package editor

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func (e *Editor) DrawDirty(s tcell.Screen, style tcell.Style) {
	w, _ := s.Size()
	for y := range e.Buffer.DirtyLine {
		l := e.Content[y]
		for x := 0; x < w; x++ {
			ch := ' '
			if l != nil && x < len(l) {
				ch = l[x]
			}
			s.SetContent(x, y-e.ScrollY, ch, nil, style)
		}
	}
	e.DirtyLine = make(map[int]struct{})
}

func (e *Editor) DrawBuffer(s tcell.Screen, style tcell.Style) {
	if len(e.Content) < e.Height {
		for y := range e.Content {
			l := e.Content[y]
			for x := 0; x < e.Width; x++ {
				ch := ' '
				if l != nil && x < len(l) {
					ch = l[x]
				}
				s.SetContent(x, y, ch, nil, style)
			}
		}

	} else {
		start := e.ScrollY
		end := min(start+e.Height, len(e.Content))
		for i := start; i < end; i++ {
			line := e.Content[i]
			y := i - start
			for x := 0; x < e.Width; x++ {
				ch := ' '
				if x < len(line) {
					ch = line[x]
				}
				s.SetContent(x, y, ch, nil, style)
			}
		}
	}
}

func (e *Editor) DrawStatusBar(s tcell.Screen, style tcell.Style) {
	w, h := s.Size()

	modeStr := "[UNKNOWN]"
	switch e.Mode {
	case Normal:
		modeStr = "[NORMAL]"
	case Insert:
		modeStr = "[INSERT]"
	}

	status := fmt.Sprintf("%s | Ln %d, Col %d lines: %d | scrollY [%d]", modeStr, e.Buffer.CursorY+1, e.Buffer.CursorX+1, len(e.Buffer.Content), e.ScrollY)
	padding := strings.Repeat(" ", max(0, w-len(status))) // fill line

	for i, r := range status + padding {
		s.SetContent(i, h-1, r, nil, style.Reverse(true)) // bottom row
	}
}
