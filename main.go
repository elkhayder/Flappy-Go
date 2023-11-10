package main

import (
	"log"

	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(shared.GameWidth, shared.GameHeight)
	ebiten.SetWindowTitle("Flappy Go!")

	game := Game{}
	game.Init()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
