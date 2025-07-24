package editor

func (e *Editor) CurLineLen() int {
	return len(e.Buffer.Content[e.Buffer.CursorY])
}
