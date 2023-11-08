package main

import (
	"bytes"
	"image"
	"log"

	AssetImages "github.com/elkhayder/Flappy-Go/assets/images"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	ebiten.SetWindowSize(288, 512)
	ebiten.SetWindowTitle("Flappy Go!")

	gameicon, _, _ := image.Decode(bytes.NewReader(AssetImages.GameIcon_png))

	var tmp []image.Image
	tmp = append(tmp, gameicon)

	ebiten.SetWindowIcon(tmp)

	var game = Game{}

	game.Init()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
