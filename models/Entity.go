package models

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	Animator Animator

	X, Y          float64
	Width, Height float64
}

func NewEntity(animator Animator, x, y float64) Entity {

	return Entity{
		Animator: animator,
		X:        x,
		Y:        y,
		// Take width and height from a single frame, we assume that all sprites have the same size
		Width:  float64(animator.CurrentFrame().Bounds().Dx()),
		Height: float64(animator.CurrentFrame().Bounds().Dy()),
	}
}

func (entity *Entity) Update(dt float64) {
	entity.Animator.Update()
}

func (entity *Entity) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	// x, y := screen.Bounds().Dx(), screen.Bounds().Dy()

	frame := entity.Animator.CurrentFrame()

	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}

	op.GeoM.Translate(
		(entity.X)-(entity.Width)/2.,
		(entity.Y)-(entity.Height)/2.,
	)

	screen.DrawImage(frame, op)
}
