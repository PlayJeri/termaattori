package editor

type Buffer struct {
	Content          [][]rune
	CursorX, CursorY int
	ScrollY          int
}
