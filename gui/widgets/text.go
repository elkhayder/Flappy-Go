package widgets

import (
	"image"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
)

type Text struct {
	color     color.Color
	horzAlign etxt.HorzAlign
	vertAlign etxt.VertAlign
	renderer  *etxt.Renderer
	size      int
}

var (
	_ furex.Drawer = (*Text)(nil)
)

func NewText(rendrer *etxt.Renderer, color color.Color, size int, horzAlign etxt.HorzAlign, vertAlign etxt.VertAlign) *Text {
	return &Text{
		renderer:  rendrer,
		color:     color,
		horzAlign: horzAlign,
		vertAlign: vertAlign,
		size:      size,
	}
}

func (t *Text) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	var (
		size int = t.size
		err  error

		vertAlign, horzAlign = t.vertAlign, t.horzAlign
	)

	if view.Attrs["size"] != "" {
		size, err = strconv.Atoi(view.Attrs["size"])
		if err != nil {
			size = t.size
		}
	}

	switch view.Attrs["align"] {
	case "center":
		horzAlign = etxt.XCenter
		vertAlign = etxt.YCenter
	}

	x, y := frame.Min.X+frame.Dx()/2, frame.Min.Y+frame.Dy()/2

	if t.horzAlign == etxt.Left {
		x = frame.Min.X
	}
	if t.vertAlign == etxt.Top {
		y = frame.Min.Y
	}

	if t.color != nil {
		t.renderer.SetColor(t.color)
	} else {
		t.renderer.SetColor(color.White)
	}

	t.renderer.SetAlign(vertAlign, horzAlign)
	t.renderer.SetTarget(screen)
	t.renderer.SetSizePx(size)

	t.renderer.Draw(view.Text, x, y)

}
