package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"strconv"

	"github.com/elkhayder/Flappy-Go/assets/sprites"
	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	bg     Parallax
	bird   Bird
	pipes  [2]PipeGroup
	hitbox CollisionBody

	score uint
	lost  bool
}

func (g *Game) Init() {
	g.bird.Init(g)

	g.bg = NewParallax(
		&color.RGBA{0x4D, 0xC1, 0xCB, 0xFF}, // Background Blue
	)

	velocity := 15.
	nextVelocity := func() float64 {
		velocity *= 1.5
		return velocity
	}

	// Should stay in the same order
	g.bg.Append(NewParallaxLayer(
		sprites.BgClouds_png,
		nextVelocity(),
		shared.GameHeight-
			shared.GroundSpriteHeight-
			shared.HillsSpriteHeight-
			shared.BushesSpriteHeight+
			10,
	))

	g.bg.Append(NewParallaxLayer(
		sprites.BgBuilding_png,
		nextVelocity(),
		shared.GameHeight-
			shared.GroundSpriteHeight-
			shared.HillsSpriteHeight-
			25,
	))

	g.bg.Append(NewParallaxLayer(
		sprites.BgBushes_png,
		nextVelocity(),
		shared.GameHeight-
			shared.GroundSpriteHeight-
			shared.HillsSpriteHeight,
	))

	g.bg.Append(NewParallaxLayer(
		sprites.BgGround_png,
		nextVelocity(),
		shared.GameHeight-
			shared.GroundSpriteHeight,
	))

	// Load PipeSprite

	pipeSpriteImg, _, err := image.Decode(bytes.NewReader(sprites.PipeRed_png))

	if err != nil {
		log.Fatal(err)
	}

	PipeSprite = ebiten.NewImageFromImage(pipeSpriteImg)

	for i := range g.pipes {
		// $velocity should now be the velocity for the pipes to match the velocity of the ground
		g.pipes[i].Init(g, velocity)
	}

	// Dont need to init min, it is auto initialized to 0,0
	g.hitbox.max.x = shared.GameWidth
	g.hitbox.max.y = shared.GameHeight - shared.GroundSpriteHeight

	g.Reset()
}

func (g *Game) Reset() {
	g.bird.Reset()

	g.lost = false
	g.score = 0
	for i := range g.pipes {
		g.pipes[i].Reset(i + 1)

	}

}

func (g *Game) Update() error {
	var birdHitBox CollisionBody

	if g.lost {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.Reset()
		}

		goto end
	}

	g.bg.Update() // Background

	g.bird.Update() // Bird
	birdHitBox = g.bird.HitBox()

	// Check if the bird is inside the screen
	if !birdHitBox.Inside(&g.hitbox) {
		g.lost = true
		goto end
	}

	for i := range g.pipes {
		pipe := &g.pipes[i]

		pipe.Update()

		top, bot := pipe.HitBox()

		// Check Collision with Bird
		if top.Overlap(&birdHitBox) || bot.Overlap(&birdHitBox) {
			g.lost = true
			goto end
		}

		// Check For Score
		if !pipe.pointCounted {
			if pipe.x < g.bird.x {
				g.score++
				pipe.pointCounted = true
			}
		} else if pipe.x < 0 && !top.Overlap(&g.hitbox) {
			// Only checking if the point is counted, because if not, we are sure it is not on the left side of screen
			// Check if it is outside the screen, dont need to check for both top and bot
			// Checking if X is less than 0 to make sure it is outlisde from the left side
			pipe.Reset(len(g.pipes)) // Move outside the screen
		}
	}

end:
	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	g.bg.Draw(screen,
		3, // index of ground sprite
	)

	// Draw Pipes
	for i := range g.pipes {
		g.pipes[i].Draw(screen)
	}

	// Draw the ground seperatly
	g.bg.DrawLayers(screen, 3)

	g.bird.Draw(screen)

	ebitenutil.DebugPrint(screen, "Score: "+fmt.Sprintf("%d", g.score)+" | Lost? "+strconv.FormatBool(g.lost))
}

func (g *Game) Layout(_, _ int) (int, int) {
	return shared.GameWidth, shared.GameHeight
}
