package editor

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func (e Editor) DrawBuffer(s tcell.Screen, style tcell.Style) {
	for row, line := range e.Buffer.Content {
		for col, r := range line {
			s.SetContent(col, row, r, nil, style)
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

	status := fmt.Sprintf("%s | Ln %d, Col %d lines: %d", modeStr, e.Buffer.CursorY+1, e.Buffer.CursorX+1, len(e.Buffer.Content))
	padding := strings.Repeat(" ", max(0, w-len(status))) // fill line

	for i, r := range status + padding {
		s.SetContent(i, h-1, r, nil, style.Reverse(true)) // bottom row
	}
}
