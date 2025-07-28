package editor

func (e *Editor) CurLineLen() int {
	return len(e.Buffer.Content[e.Buffer.CursorY])
}

func (e *Editor) GetCurLine() []rune {
	return e.Buffer.Content[e.Buffer.CursorY]
}

func (e *Editor) MoveCursorY(dir string) {
	switch dir {
	case "up":
		e.CursorY--
		e.SetScrollY()
	case "down":
		e.CursorY++
		e.SetScrollY()
	default:
		return
	}
}

func (e *Editor) SetScrollY() {
	if e.CursorY < e.ScrollY {
		e.ScrollY = e.CursorY
	} else if e.CursorY >= e.ScrollY+e.Height {
		e.ScrollY = e.CursorY - e.Height + 1
	}
}
