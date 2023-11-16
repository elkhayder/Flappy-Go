package game

import (
	"github.com/elkhayder/Flappy-Go/assets/sprites"
	"github.com/elkhayder/Flappy-Go/config"
	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Bird struct {
	game     *Game
	animator Animator

	x, y float64

	velocity float64 // px per sec
}

func (b *Bird) Init(g *Game) {
	b.animator = NewAnimator(
		100,
		&sprites.BirdDownflap_png,
		&sprites.BirdMidflap_png,
		&sprites.BirdUpflap_png,
		&sprites.BirdMidflap_png,
	)

	b.game = g

	b.Reset()
}

func (b *Bird) Reset() {
	b.x = shared.GameWidth / 3
	b.y = shared.GameHeight / 2
	b.velocity = 0
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

	width, height := b.animator.Bounds().Dx(), b.animator.Bounds().Dy()

	frame := b.animator.CurrentFrame()

	op.GeoM.Translate(
		(b.x)-float64(width)/2.,
		(b.y)-float64(height)/2.,
	)

	screen.DrawImage(frame, &op)

}

func (b *Bird) Update() {
	// b.animator.SetFrameDuration(-100.0/shared.Gravity*math.Abs(b.velocity) + 200.)

	b.animator.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {

		b.velocity = -shared.Gravity * 0.8 // Experimental value

		// Play FX
		b.game.soundManager.PlayFx(config.FxJump)
	}

	b.velocity += shared.Gravity * shared.Dt

	b.y += b.velocity * shared.Dt
}

func (b *Bird) HitBox() CollisionBody {
	width, height := b.animator.Bounds().Dx(), b.animator.Bounds().Dy()

	base := CollisionBody{
		rectangles: []Rectangle{
			{min: Point{12, 0}, max: Point{20, 24}},
			{min: Point{8, 2}, max: Point{26, 22}},
			{min: Point{2, 8}, max: Point{32, 20}},
		},
		outer: Rectangle{
			min: Point{0, 0},
			max: Point{float64(width), float64(height)},
		},
	}

	base.CenterAround(b.x, b.y)

	return base
}
