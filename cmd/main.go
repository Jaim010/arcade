package main

import (
	"github.com/Jaim010/hangman/pkg/game"
	"github.com/Jaim010/hangman/pkg/ui"
)

func main() {
	ui.Init()
	game := game.New()
	game.Run()
}
