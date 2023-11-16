package main

import (
	"bytes"
	"image"
	"log"

	image_assets "github.com/elkhayder/Flappy-Go/assets/images"
	"github.com/elkhayder/Flappy-Go/config"
	"github.com/elkhayder/Flappy-Go/game"
	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mat/besticon/ico"
)

func main() {

	ebiten.SetWindowSize(shared.GameWidth, shared.GameHeight)
	ebiten.SetWindowTitle("Flappy Go!")
	ebiten.SetTPS(shared.TPS)
	ebiten.SetVsyncEnabled(false)

	defer config.SaveConfig()

	favicon, err := ico.Decode(bytes.NewReader(image_assets.Favicon_ico))
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowIcon([]image.Image{favicon})

	game := game.Game{}
	game.Init()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}

}
