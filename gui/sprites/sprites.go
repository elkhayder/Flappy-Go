package gui_sprites

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var sprites = make(map[string]*ebiten.Image)

//go:embed assets/*
var fs embed.FS

func Load() {
	dir, err := fs.ReadDir("assets")

	if err != nil {
		log.Fatal(err)
	}

	for i := range dir {
		name := dir[i].Name()
		raw, err := fs.ReadFile("assets/" + name)

		if err != nil {
			log.Fatal(err)
		}

		image, _, err := image.Decode(bytes.NewReader(raw))

		if err != nil {
			log.Fatal(err)
		}

		sprites[name] = ebiten.NewImageFromImage(image)
	}

	log.Println("GUI assets loaded")
}

func Get(name string) *ebiten.Image {
	sprite := sprites[name]

	if sprite == nil {
		log.Fatalf("Sprite %s not found", name)
	}

	return sprite
}
