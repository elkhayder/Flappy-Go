package widgets

import (
	"image"

	gui_sprites "github.com/elkhayder/Flappy-Go/gui/sprites"
	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
)

type Card struct {
}

var _ furex.Drawer = (*Card)(nil)

func (c *Card) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	var (
		Header = gui_sprites.Get("Card_Header.png")
		Body   = gui_sprites.Get("Card_Body.png")
		Footer = gui_sprites.Get("Card_Footer.png")

		XOrigin               = float64(shared.GameWidth-Header.Bounds().Dx()) / 2
		DiffHeaderBodyOneSide = float64(Header.Bounds().Dx()-Body.Bounds().Dx()) / 2
	)

	op := ebiten.DrawImageOptions{}
	op.GeoM.Reset()

	// Draw Header
	op.GeoM.Translate(XOrigin, float64(frame.Min.Y))
	screen.DrawImage(Header, &op)

	// Draw Body
	// Each BodyImage is a single Pixel in height
	op.GeoM.Translate(DiffHeaderBodyOneSide, float64(Header.Bounds().Dy()))

	DrawnHeight := Header.Bounds().Dy()

	for DrawnHeight <= frame.Dy() {
		screen.DrawImage(Body, &op)
		op.GeoM.Translate(0, 1)
		DrawnHeight++
	}

	// Draw Footer
	screen.DrawImage(Footer, &op)
}
