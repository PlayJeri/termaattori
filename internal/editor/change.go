package editor

type Snapshot struct {
	Buf [][]rune
}

func (b *Buffer) AppendUndo() {
	old := b.CopyBuffer()
	b.UndoStack = append(b.UndoStack, Snapshot{Buf: old})
	b.RedoStack = []Snapshot{}
}

func (b *Buffer) SaveSnapshot() {
	if len(b.UndoStack) == 0 {
		b.AppendUndo()
		return
	}
	if len(b.UndoStack[len(b.UndoStack)-1].Buf) != len(b.Content) {
		old := b.CopyBuffer()
		b.UndoStack = append(b.UndoStack, Snapshot{Buf: old})
		return
	}
	for i := range b.Content {
		if !equalRunes(b.Content[i], b.UndoStack[len(b.UndoStack)-1].Buf[i]) {
			b.AppendUndo()
			return
		}
	}
}

func (b *Buffer) CopyBuffer() [][]rune {
	content := make([][]rune, len(b.Content))

	for i := range content {
		content[i] = append([]rune(nil), b.Content[i]...)
	}

	return content
}

func (b *Buffer) Undo() {
	if len(b.UndoStack) == 0 {
		return
	}

	stateToReturn := b.UndoStack[len(b.UndoStack)-1].Buf
	b.UndoStack = b.UndoStack[:len(b.UndoStack)-1]
	current := b.CopyBuffer()
	b.RedoStack = append(b.RedoStack, Snapshot{Buf: current})

	if len(stateToReturn) < len(b.Content) {
		b.Content = b.Content[:len(stateToReturn)]
	}
	for i := range stateToReturn {
		b.Content[i] = append([]rune(nil), stateToReturn[i]...)
	}

	b.CheckCursor()
}

func (b *Buffer) Redo() {
	if len(b.RedoStack) == 0 {
		return
	}

	stateToReturn := b.RedoStack[len(b.RedoStack)-1].Buf
	b.RedoStack = b.RedoStack[:len(b.RedoStack)-1]
	current := b.CopyBuffer()
	b.UndoStack = append(b.UndoStack, Snapshot{Buf: current})

	if len(stateToReturn) < len(b.Content) {
		b.Content = b.Content[:len(stateToReturn)]
	}

	for i := range stateToReturn {
		b.Content[i] = append([]rune(nil), stateToReturn[i]...)
	}

	b.CheckCursor()
}

func (b *Buffer) CheckCursor() {
	if b.CursorY >= len(b.Content) {
		b.CursorY = len(b.Content) - 1
	}

	if b.CursorX > len(b.Content[b.CursorY]) {
		b.CursorX = len(b.Content[b.CursorY])
	}
}
