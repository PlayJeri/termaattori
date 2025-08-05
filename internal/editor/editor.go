package editor

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

// Constants
const (
	GutterWidth     = 4
	StatusBarHeight = 1
)

type Mode int

const (
	Normal Mode = iota
	Insert
	Visual
)

type Editor struct {
	Buffer
	Mode          Mode
	Width, Height int
}

func NewEditor() *Editor {
	return &Editor{
		Buffer: Buffer{
			Content:   [][]rune{{}},
			CursorX:   0,
			CursorY:   0,
			DirtyLine: make(map[int]struct{}),
			ScrollY:   0,
		},
		Mode: Normal,
	}
}

func (e *Editor) Run(s tcell.Screen, style tcell.Style) {
	defer s.Fini()
	e.Width, e.Height = s.Size()
	e.DrawBuffer(s, style)
	for {
		// Update cursor
		switch e.Mode {
		case Normal:
			s.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
		case Insert:
			s.SetCursorStyle(tcell.CursorStyleBlinkingBar)
		}
		s.ShowCursor(e.Buffer.CursorX+GutterWidth, e.CursorY-e.ScrollY)

		// Update screen
		if len(e.DirtyLine) > 0 {
			e.DrawDirty(s, style)
		}
		e.DrawGutter(s, style)
		e.DrawStatusBar(s, style)
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				s.Fini()
				os.Exit(0)
			}
			switch e.Mode {
			case Normal:
				e.handleNormalMode(ev)
			case Insert:
				e.handleInsertMode(ev)
			}
		}
	}

}
