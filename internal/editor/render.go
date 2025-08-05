package editor

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func (e *Editor) DrawDirty(s tcell.Screen, style tcell.Style) {
	w := e.Width
	w -= GutterWidth
	for y := range e.Buffer.DirtyLine {
		l := e.Content[y]
		for x := 0; x < w; x++ {
			ch := ' '
			if l != nil && x < len(l) {
				ch = l[x]
			}
			s.SetContent(x+GutterWidth, y-e.ScrollY, ch, nil, style)
		}
	}
	e.DirtyLine = make(map[int]struct{})
}

func (e *Editor) DrawBuffer(s tcell.Screen, style tcell.Style) {
	// bufXStart := GutterWidth
	// bufYStart := StatusBarHeight
	bufWidth := e.Width - GutterWidth
	bufHeight := e.Height - StatusBarHeight

	start := e.ScrollY
	end := min(start+bufHeight, len(e.Content))
	for i := start; i < end; i++ {
		line := e.Content[i]
		y := i - start
		for x := 0; x < bufWidth; x++ {
			ch := ' '
			if x < len(line) {
				ch = line[x]
			}
			s.SetContent(x+GutterWidth, y, ch, nil, style)
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

	status := fmt.Sprintf("%s | Ln %d, Col %d", modeStr, e.Buffer.CursorY+1, e.Buffer.CursorX+1)
	padding := strings.Repeat(" ", max(0, w-len(status))) // fill line

	for i, r := range status + padding {
		s.SetContent(i, h-1, r, nil, style.Reverse(true)) // bottom row
	}
}

func (e *Editor) DrawGutter(s tcell.Screen, style tcell.Style) {
	for i := 0; i < e.Height-1; i++ {
		lineNum := fmt.Sprintf("%*d ", GutterWidth-1, i+1)
		for j, ch := range lineNum {
			s.SetContent(j, i, ch, nil, style)
		}
	}
}
