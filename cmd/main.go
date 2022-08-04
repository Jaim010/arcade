package main

import (
	"github.com/Jaim010/arcade/pkg/hangman/game"
	"github.com/Jaim010/arcade/pkg/utils/screen"
)

func main() {
	screen.Init()
	game := game.New()
	game.Run()
}
