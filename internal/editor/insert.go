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
		e.BufferChanged = true
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
		e.BufferChanged = true
		e.AddLine()
	case tcell.KeyEsc:
		e.ChangeMode(Normal)
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		e.BufferChanged = true
		e.RemoveChar()
	}
}

func (e *Editor) RemoveChar() {
	if e.CursorX > 0 {
		line := e.GetCurLine()
		line = slices.Delete(line, e.CursorX-1, e.CursorX)
		e.Content[e.CursorY] = line
		e.CursorX--
	} else {
		e.RemoveLine()
	}
}

func (e *Editor) RemoveLine() {
	if e.CursorY == 0 {
		return
	}
	curY := e.CursorY
	lineToDelete := e.GetCurLine()
	e.MoveCursorY("up")
	lineToJoin := e.GetCurLine()
	e.CursorX = len(lineToJoin)
	e.Content[e.CursorY] = append(lineToJoin, lineToDelete...)
	e.Content = slices.Delete(e.Content, curY, curY+1)
}

func (e *Editor) AddLine() {
	if len(e.Content) == 0 {
		e.Content = append(e.Content, []rune{})
		e.CursorX = 0
		e.MoveCursorY("down")
		return
	}
	e.Content = slices.Insert(e.Content, e.CursorY+1, []rune{})
	e.CursorX = 0
	e.MoveCursorY("down")
}

func (e *Editor) handleInsertRune(r rune) {
	if len(e.Content) == 0 {
		e.Content = append(e.Content, []rune{})
	}
	line := e.Content[e.CursorY]
	line = append(line[:e.CursorX], append([]rune{r}, line[e.CursorX:]...)...)
	e.Content[e.CursorY] = line
	e.CursorX++
}
