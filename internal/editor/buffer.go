package editor

type Buffer struct {
	Content          [][]rune
	CursorX, CursorY int
	ScrollY          int
	UndoStack        []Snapshot
	RedoStack        []Snapshot
	BufferChanged    bool
}
