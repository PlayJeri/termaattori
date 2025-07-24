package editor

import (
	"slices"

	"github.com/gdamore/tcell/v2"
)

func (e *Editor) handleInsertMode(ev *tcell.EventKey) {
	r, k, m := ev.Rune(), ev.Key(), ev.Modifiers()

	if m != 0 {
		//TODO: add modifier handler
		return
	}

	if k == tcell.KeyRune {
		e.handleInsertRune(r)
		return
	}

	if k != 0 {
		e.handleNormalModeKey(k)
		return
	}
}

func (e *Editor) handleNormalModeKey(k tcell.Key) {
	switch k {
	case tcell.KeyEnter:
		e.AddLine()
	case tcell.KeyEsc:
		e.Mode = Normal
	}
}

func (e *Editor) AddLine() {
	e.Buffer.Content = slices.Insert(e.Buffer.Content, e.Buffer.CursorY+1, []rune{})
	e.Buffer.CursorX = 0
	e.Buffer.CursorY++
	e.NeedsFullRedraw = true
}

func (e *Editor) handleInsertRune(r rune) {
	if len(e.Buffer.Content) == 0 {
		e.Buffer.Content = append(e.Buffer.Content, []rune{})
	}
	line := e.Buffer.Content[e.Buffer.CursorY]
	line = append(line[:e.Buffer.CursorX], append([]rune{r}, line[e.Buffer.CursorX:]...)...)
	e.Buffer.Content[e.Buffer.CursorY] = line
	e.Buffer.CursorX++
}
