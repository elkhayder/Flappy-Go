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

//go:embed assets/**
var fs embed.FS

func Load() {
	parseDirectory("")

	log.Println("GUI assets loaded")
}

func parseDirectory(path string) {
	// folder := "assets"

	// if path != "" {
	// 	folder += "/" + path
	// }

	dir, err := fs.ReadDir("assets" + path)

	if err != nil {
		log.Fatal(err)
	}

	for i := range dir {
		name := dir[i].Name()

		if dir[i].IsDir() {
			parseDirectory(path + "/" + name)
		} else {
			raw, err := fs.ReadFile("assets" + path + "/" + name)

			if err != nil {
				log.Fatal(err)
			}

			image, _, err := image.Decode(bytes.NewReader(raw))

			if err != nil {
				log.Fatal(err)
			}

			cleanpath := path

			if len(path) > 0 {
				cleanpath = cleanpath[1:]
			}

			sprites[cleanpath+"/"+name] = ebiten.NewImageFromImage(image)
		}
	}
}

func Get(name string) *ebiten.Image {
	sprite := sprites[name]

	if sprite == nil {
		log.Fatalf("Sprite %s not found", name)
	}

	return sprite
}
