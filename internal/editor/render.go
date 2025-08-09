package editor

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func (e *Editor) DrawBuffer(s tcell.Screen, style tcell.Style) {
	bufWidth := e.Width - GutterWidth
	bufHeight := e.Height - StatusBarHeight

	start := e.ScrollY
	// end := min(start+bufHeight, len(e.Content))
	end := start + bufHeight
	for i := start; i < end; i++ {
		var line []rune
		if i < len(e.Content) {
			line = e.Content[i]
		}
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
		if i+e.ScrollY < len(e.Content) {
			ln := i + 1 + e.ScrollY
			lineNum := fmt.Sprintf("%*d ", GutterWidth-1, ln)
			for j, ch := range lineNum {
				s.SetContent(j, i, ch, nil, style)
			}
		} else {
			for j := range GutterWidth {
				s.SetContent(j, i, ' ', nil, style)
			}
		}
	}
}
