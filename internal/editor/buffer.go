package editor

type Buffer struct {
	Content          [][]rune
	CursorX, CursorY int
	DirtyLine        map[int]struct{}
	ScrollY          int
}
