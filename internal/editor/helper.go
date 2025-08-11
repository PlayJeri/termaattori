package editor

func (e *Editor) CurLineLen() int {
	return len(e.Buffer.Content[e.Buffer.CursorY])
}

func (e *Editor) GetCurLine() []rune {
	line := e.Buffer.Content[e.Buffer.CursorY]
	c := make([]rune, len(line))
	copy(c, line)
	return c
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
	visibleLines := max(e.Height-1, 1)

	if e.CursorY < e.ScrollY {
		e.ScrollY = e.CursorY
	} else if e.CursorY >= e.ScrollY+visibleLines {
		maxScroll := max(len(e.Content)-visibleLines, 0)
		e.ScrollY = min(e.CursorY-visibleLines+1, maxScroll)
	}
}

func equalRunes(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
