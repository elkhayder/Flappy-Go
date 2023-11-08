package entities

import (
	"math"
	"math/rand"

	PipeSprites "github.com/elkhayder/Flappy-Go/assets/sprites/pipe"
	"github.com/elkhayder/Flappy-Go/models"
)

const (
	PipeGroupVelocity = 60  // px/sec
	PipeGroupGap      = 100 // px
)

type Pipe struct {
	models.Entity

	ScoreCounted bool
}

func NewPipe(x, y float64) Pipe {
	pipe := Pipe{
		Entity: models.NewEntity(
			models.NewAnimator(
				0,
				&PipeSprites.Green,
			),
			x, y,
		),
	}

	return pipe
}

func (b *Pipe) Update(dt float64) {
	b.Entity.Update(dt)

	// divide by 1000 to go from ms to s

	b.Entity.X -= PipeGroupVelocity * dt / 1e3
}

func (p *Pipe) RandomizePositionOutsideRightScreen(width, height int) {
	p.X = float64(width) + p.Width/2

	p.Y = float64(rand.Intn(height-PipeGroupGap/2) + PipeGroupGap/2)
}

func (p *Pipe) CheckCollision(b *Bird) bool {
	isInsideThePipeVertically := b.X+b.Width/2 > p.X-p.Width/2 && b.X-b.Width/2 < p.X+p.Width/2

	if !isInsideThePipeVertically {
		return false
	}

	HorizontalCollision := math.Abs(p.Y-b.Y)+b.Height/2 > PipeGroupGap/2

	return HorizontalCollision
}

func (p *Pipe) IsBehindTheBird(b *Bird) bool {
	return p.X+p.Width/2 < b.X-b.Width/2
}

func (p *Pipe) IsOutsideScreenLeft() bool {
	return p.X+p.Width/2 < 0
}
