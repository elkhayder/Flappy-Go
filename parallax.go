package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"

	"log"

	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

/// Single Parallax Layer

/*
 * For now, only horizontal parallax is supported (x axis)
 */
type ParallaxLayer struct {
	sprite *ebiten.Image // original Sprite

	velocity float64 // px per second
	offset   float64 // px
	y        int     // Vertical position - Constant
}

func NewParallaxLayer(imgRaw []byte, velocity float64, y int) ParallaxLayer {
	img, _, err := image.Decode(bytes.NewReader(imgRaw))

	if err != nil {
		log.Fatal(err)
	}

	sprite := ebiten.NewImageFromImage(img)

	return ParallaxLayer{
		velocity: velocity,
		y:        y,

		sprite: sprite,

		// image: ebiten.NewImage(shared.GameWidth, sprite.Bounds().Dy()),
	}
}

func (l *ParallaxLayer) Update() {
	// Update offset
	dx := shared.Dt * l.velocity

	l.offset += dx

	if l.offset > float64(l.sprite.Bounds().Dx()) { // Reset to 0
		l.offset = 0
	}

}

func (l *ParallaxLayer) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}

	drawOffset := -l.offset

	for drawOffset <= shared.GameWidth {
		op.GeoM.Reset()
		op.GeoM.Translate(float64(drawOffset), float64(l.y))
		screen.DrawImage(l.sprite, &op)

		drawOffset += float64(l.sprite.Bounds().Dx())
	}
}

/// Parallax Group

type Parallax struct {
	layers []ParallaxLayer
	fill   *color.RGBA
}

func NewParallax(fill *color.RGBA) Parallax {
	return Parallax{
		layers: make([]ParallaxLayer, 0),
		fill:   fill,
	}
}

func (p *Parallax) Append(layer ParallaxLayer) {
	p.layers = append(p.layers, layer)
}

func (p *Parallax) Update() {
	for i := range p.layers {
		p.layers[i].Update()
	}
}

func (p *Parallax) Draw(screen *ebiten.Image) {
	if p.fill != nil {
		screen.Fill(*p.fill)
	}

	for i := range p.layers {
		p.layers[i].Draw(screen)
	}
}
