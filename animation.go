package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"

	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

type Animator struct {
	// Sprites
	sprites            []*ebiten.Image
	spritesCount       int
	currentSpriteIndex int

	// Management
	ticksPerFrame int64
	ticksCount    int64
}

func NewAnimator(MillisecondsPerFrame float64, rawSprites ...*[]byte) Animator {
	// Parse sprites
	sprites := make([]*ebiten.Image, len(rawSprites))

	for i, raw := range rawSprites {
		img, _, err := image.Decode(bytes.NewReader(*raw))

		if err != nil {
			log.Fatal(err)
		}

		sprites[i] = ebiten.NewImageFromImage(img)
	}

	animtor := Animator{
		sprites:            sprites,
		spritesCount:       len(sprites),
		currentSpriteIndex: 0,
		//
		ticksPerFrame: 0,
	}

	animtor.SetFrameDuration(MillisecondsPerFrame)

	return animtor
}

func (a *Animator) SetFrameDuration(ms float64) {
	a.ticksPerFrame = int64(ms / 1000.0 * shared.TPS)
}

func (a *Animator) CurrentFrame() *ebiten.Image {
	return a.sprites[a.currentSpriteIndex]
}

func (a *Animator) Bounds() image.Rectangle {
	return a.CurrentFrame().Bounds()
}

func (a *Animator) Update() {
	// Static Animation (Single frame)
	if a.ticksPerFrame == 0 {
		return
	}

	a.ticksCount += 1

	if a.ticksCount > a.ticksPerFrame {
		a.currentSpriteIndex = (a.currentSpriteIndex + 1) % a.spritesCount
		a.ticksCount = 0
	}
}
