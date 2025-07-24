package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/playjeri/termaattori/internal/editor"
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	e := editor.NewEditor()

	e.Run(s, defStyle)
}
