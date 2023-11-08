package entities

import (
	"time"

	BirdSprites "github.com/elkhayder/Flappy-Go/assets/sprites/bird"
	"github.com/elkhayder/Flappy-Go/models"
	"github.com/hajimehoshi/ebiten/v2"
)

const Gravity = 200 // px/sec^2

type Bird struct {
	models.Entity
	velocity float64 // px per sec

	lastJumpAt int64 // Last Jump
}

func (b *Bird) Init(x, y float64) {
	b.Entity = models.NewEntity(
		models.NewAnimator(
			100,
			&BirdSprites.Downflap_png,
			&BirdSprites.Midflap_png,
			&BirdSprites.Upflap_png,
			&BirdSprites.Midflap_png,
		),
		x,
		y,
	)
}

func (b *Bird) Draw(screen *ebiten.Image) {

	op := ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	// // Calculate rotation angle based on velocity
	//TODO: Calculate rotation angle
	// b.angle += 60. / 200. * b.velocity

	// if b.angle > 30 {
	// 	b.angle = 30
	// } else if b.angle < -30 {
	// 	b.angle = -30
	// }

	// op.GeoM.Rotate(b.angle * 2 * math.Pi / 360)

	b.Entity.Draw(screen, &op)
}

func (b *Bird) Update(dt float64) {
	b.Entity.Update(dt)

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		b.handleJump(dt)
	}

	// divide by 1000 to go from ms to s

	b.velocity += Gravity * dt / 1e3

	b.Entity.Y += b.velocity * dt / 1e3
}

func (b *Bird) handleJump(dt float64) {
	const JumpDelay = 500 // ms

	now := time.Now().UnixMilli()

	if now-b.lastJumpAt > JumpDelay {
		b.lastJumpAt = now

		b.velocity = -Gravity * .8
	}
}

func (b *Bird) CheckInsideScreen(width, height int) bool {
	return b.Entity.X-b.Entity.Width/2 > 0 &&
		b.Entity.X+b.Entity.Width/2 <= float64(width) &&
		b.Entity.Y-b.Entity.Height/2 > 0 &&
		b.Entity.Y+b.Entity.Height/2 <= float64(height)
}
