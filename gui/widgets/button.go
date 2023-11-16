package widgets

import (
	"image"
	"image/color"

	gui_sprites "github.com/elkhayder/Flappy-Go/gui/sprites"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
)

type Button struct {
	Color   color.Color
	OnClick func()
	Text    *etxt.Renderer

	mouseover bool
	pressed   bool
}

var (
	_ furex.ButtonHandler          = (*Button)(nil)
	_ furex.Drawer                 = (*Button)(nil)
	_ furex.MouseEnterLeaveHandler = (*Button)(nil)
)

func (b *Button) HandlePress(x, y int, t ebiten.TouchID) {
	b.pressed = true
}

func (b *Button) HandleRelease(x, y int, isCancel bool) {
	b.pressed = false
	if !isCancel {
		if b.OnClick != nil {
			b.OnClick()
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	centerX, centerY := frame.Min.X+frame.Dx()/2, frame.Min.Y+frame.Dy()/2

	var sprite *ebiten.Image

	if b.pressed {
		sprite = gui_sprites.Get(view.Attrs["sprite_pressed"])
	} else {
		sprite = gui_sprites.Get(view.Attrs["sprite"])
	}

	scaleX, scaleY := float64(frame.Dx())/float64(sprite.Bounds().Dx()), float64(frame.Dy())/float64(sprite.Bounds().Dy())

	op := ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	op.GeoM.Scale(scaleX, scaleY)

	screen.DrawImage(sprite, &op)

	b.Text.SetAlign(etxt.YCenter, etxt.XCenter)
	b.Text.SetTarget(screen)
	if b.Color != nil {
		b.Text.SetColor(b.Color)
	} else {
		b.Text.SetColor(color.White)
	}
	b.Text.Draw(view.Text, centerX, centerY)
}

func (b *Button) HandleMouseEnter(x, y int) bool {
	b.mouseover = true
	return true
}

func (b *Button) HandleMouseLeave() {
	b.mouseover = false
}
