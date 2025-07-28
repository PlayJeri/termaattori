package editor

import "github.com/gdamore/tcell/v2"

func (e *Editor) handleNormalMode(ev *tcell.EventKey) {
	r, k, m := ev.Rune(), ev.Key(), ev.Modifiers()

	if m != 0 {
		//TODO: add modifier handler
		return
	}

	if k == tcell.KeyRune {
		e.handleNormalModeRune(r)
		return
	}

	if k != 0 {
		// e.handleNormalModeKey(k)
		return
	}
}

func (e *Editor) handleNormalModeRune(r rune) {
	switch r {
	case 'i':
		e.Mode = Insert
	case 'a':
		e.Mode = Insert
		if e.Buffer.CursorX < e.CurLineLen() {
			e.Buffer.CursorX++
		}
	case 'h':
		e.MoveLeft()
	case 'j':
		e.MoveDown()
	case 'k':
		e.MoveUp()
	case 'l':
		e.MoveRight()
	}

}

func (e *Editor) MoveLeft() {
	if e.Buffer.CursorX > 0 {
		e.Buffer.CursorX--
	}
}

func (e *Editor) MoveRight() {
	if e.Buffer.CursorX < e.CurLineLen() {
		e.Buffer.CursorX++
	}
}

func (e *Editor) MoveUp() {
	if e.Buffer.CursorY > 0 {
		e.MoveCursorY("up")
		e.Buffer.CursorX = 0
	}
}

func (e *Editor) MoveDown() {
	if e.Buffer.CursorY+1 < len(e.Buffer.Content) {
		e.MoveCursorY("down")
		e.Buffer.CursorX = 0
	}
}
