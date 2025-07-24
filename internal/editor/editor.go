package editor

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

type Mode int

const (
	Normal Mode = iota
	Insert
	Visual
)

type Editor struct {
	Buffer          Buffer
	Mode            Mode
	NeedsFullRedraw bool
}

func NewEditor() *Editor {
	return &Editor{
		Buffer: Buffer{
			Content: [][]rune{},
			CursorX: 0,
			CursorY: 0,
		},
		Mode:            Normal,
		NeedsFullRedraw: false,
	}
}

func (e *Editor) Run(s tcell.Screen, style tcell.Style) {
	defer s.Fini()
	for {
		// Update cursor
		switch e.Mode {
		case Normal:
			s.SetCursorStyle(tcell.CursorStyleBlinkingBlock)
		case Insert:
			s.SetCursorStyle(tcell.CursorStyleBlinkingBar)
		}
		s.ShowCursor(e.Buffer.CursorX, e.Buffer.CursorY)

		// Update screen
		e.DrawBuffer(s, style)
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

		if e.NeedsFullRedraw {
			s.Clear()
			e.NeedsFullRedraw = false
		}
	}

}
