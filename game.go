package main

import (
	"fmt"
	"image/color"

	sprites "github.com/elkhayder/Flappy-Go/assets/sprites/background"
	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	bg Parallax
}

func (g *Game) Init() {
	g.bg = NewParallax(
		&color.RGBA{0x4D, 0xC1, 0xCB, 0xFF}, // Background Blue
	)

	const (
		GroundHeight = 112
		BushesHeight = 66
		HillsHeight  = 34
	)

	velocity := 10.
	nextVelocity := func() float64 {
		velocity *= 1.7
		return velocity
	}

	// Should stay in the same order
	g.bg.Append(NewParallaxLayer(
		sprites.BgClouds_png,
		nextVelocity(),
		shared.GameHeight-GroundHeight-HillsHeight-BushesHeight+10,
	))

	g.bg.Append(NewParallaxLayer(
		sprites.BgBuilding_png,
		nextVelocity(),
		shared.GameHeight-GroundHeight-HillsHeight-25,
	))

	g.bg.Append(NewParallaxLayer(
		sprites.BgBushes_png,
		nextVelocity(),
		shared.GameHeight-GroundHeight-HillsHeight,
	))

	g.bg.Append(NewParallaxLayer(
		sprites.BgGround_png,
		nextVelocity(),
		shared.GameHeight-GroundHeight,
	))

	// $velocity should now be the velocity for the pipes to match the velocity of the ground
}

func (g *Game) Update() error {
	g.bg.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.bg.Draw(screen)

	ebitenutil.DebugPrint(screen, "FPS: "+fmt.Sprintf("%f", ebiten.ActualFPS()))
}

func (g *Game) Layout(_, _ int) (int, int) {
	return shared.GameWidth, shared.GameHeight
}
