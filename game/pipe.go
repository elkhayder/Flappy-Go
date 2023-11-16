package game

import (
	"log"
	"math/rand"

	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PipeGroupGap  = 100
	PipeGroupYmin = PipeGroupGap/2 + 30
	PipeGroupYmax = shared.GameHeight - shared.GroundSpriteHeight - PipeGroupYmin
)

var PipeSprite *ebiten.Image

type PipeGroup struct {
	game         *Game
	x, y         float64
	velocity     float64 // px/s
	pointCounted bool
}

func (g *PipeGroup) Init(game *Game, velocity float64) {
	g.game = game
	g.velocity = velocity
}

func (g *PipeGroup) Reset(offset int) {
	g.y = PipeGroupYmin + float64(rand.Intn(PipeGroupYmax-PipeGroupYmin))
	g.x = 3 * float64(shared.GameWidth) / 4 * float64(offset)
	g.pointCounted = false
}

func (g *PipeGroup) Update() {
	g.x -= g.velocity * shared.Dt
}

func (g *PipeGroup) Draw(screen *ebiten.Image) {
	if PipeSprite == nil {
		log.Fatal("PipeSprite is nil")
	}

	// Bottom Pipe
	op := ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	op.GeoM.Translate(
		g.x-float64(shared.PipeSpriteWidth)/2,
		g.y+PipeGroupGap/2,
	)

	screen.DrawImage(PipeSprite, &op)

	// Top Pipe
	op.GeoM.Reset()

	op.GeoM.Scale(1, -1)

	op.GeoM.Translate(
		g.x-float64(shared.PipeSpriteWidth)/2,
		g.y-PipeGroupGap/2,
	)

	screen.DrawImage(PipeSprite, &op)

}

func (g *PipeGroup) HitBox() CollisionBody { // return Top and Bot hitboxes
	base := CollisionBody{
		outer: Rectangle{
			min: Point{0, 0},
			max: Point{shared.PipeSpriteWidth, shared.PipeSpriteHeight*2 + PipeGroupGap},
		},
		rectangles: []Rectangle{
			{min: Point{0, 0}, max: Point{shared.PipeSpriteWidth, shared.PipeSpriteHeight}},
			{
				min: Point{0, shared.PipeSpriteHeight + PipeGroupGap},
				max: Point{shared.PipeSpriteWidth, 2*shared.PipeSpriteHeight + PipeGroupGap},
			},
		},
	}

	base.CenterAround(g.x, g.y)

	return base
}
