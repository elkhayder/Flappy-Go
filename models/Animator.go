package models

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animator struct {
	// Sprites
	sprites            []*ebiten.Image
	spritesCount       int
	currentSpriteIndex int

	// Management
	MillisecondsPerFrame int64 // in Seconds
	lastUpdatedAt        int64
}

func NewAnimator(MillisecondsPerFrame int64, rawSprites ...*[]byte) Animator {
	// Parse sprites
	sprites := make([]*ebiten.Image, len(rawSprites))

	for i, raw := range rawSprites {
		img, _, err := image.Decode(bytes.NewReader(*raw))

		if err != nil {
			log.Fatal(err)
		}

		sprites[i] = ebiten.NewImageFromImage(img)
	}

	return Animator{
		sprites:            sprites,
		spritesCount:       len(sprites),
		currentSpriteIndex: 0,
		//
		MillisecondsPerFrame: MillisecondsPerFrame,
	}

}

func (animator *Animator) CurrentFrame() *ebiten.Image {
	return animator.sprites[animator.currentSpriteIndex]
}

func (animator *Animator) Update() {
	// Static Animation (Single frame)
	if animator.MillisecondsPerFrame == 0 {
		return
	}

	now := time.Now().UnixMilli()

	if now-animator.lastUpdatedAt > animator.MillisecondsPerFrame {
		// Increment the current Sprite Index, and wrap around if it surpasses count
		animator.currentSpriteIndex = (animator.currentSpriteIndex + 1) % animator.spritesCount
		// Save the current time
		animator.lastUpdatedAt = now
	}
}
