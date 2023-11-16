package widgets

import (
	"image"

	gui_sprites "github.com/elkhayder/Flappy-Go/gui/sprites"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
)

type Checkbox struct {
	pressed bool
	model   *bool
}

var (
	_ furex.ButtonHandler = (*Checkbox)(nil)
	_ furex.Drawer        = (*Checkbox)(nil)
)

func NewCheckbox(model *bool) *Checkbox {
	return &Checkbox{
		model:   model,
		pressed: false,
	}
}

func (c *Checkbox) HandlePress(x, y int, t ebiten.TouchID) {
	c.pressed = true
}

func (c *Checkbox) HandleRelease(x, y int, isCancel bool) {
	c.pressed = false

	if !isCancel {
		*c.model = !*c.model
	}
}

func (c *Checkbox) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	var sprite *ebiten.Image

	if *c.model {
		sprite = gui_sprites.Get("Checkbox_Checked.png")
	} else {
		sprite = gui_sprites.Get("Checkbox_Unchecked.png")
	}

	scaleX, scaleY := float64(frame.Dx())/float64(sprite.Bounds().Dx()), float64(frame.Dy())/float64(sprite.Bounds().Dy())

	op := ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	op.GeoM.Scale(scaleX, scaleY)

	op.GeoM.Translate(
		float64(frame.Min.X),
		float64(frame.Min.Y),
	)

	screen.DrawImage(sprite, &op)

}
