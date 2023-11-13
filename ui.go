package main

import (
	"bytes"
	"image"
	"image/color"
	"log"

	"github.com/elkhayder/Flappy-Go/assets/fonts"
	image_assets "github.com/elkhayder/Flappy-Go/assets/images"
	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/opentype"

	"github.com/tinne26/etxt"
)

type UI struct {
	digits [10]*ebiten.Image
	// game   *main.Game

	text *etxt.Renderer
}

func (ui *UI) Init(
// game *main.Game
) {
	// ui.game = game

	// Load Digits
	for i, raw := range [10]*[]byte{
		&image_assets.Digit0_png,
		&image_assets.Digit1_png,
		&image_assets.Digit2_png,
		&image_assets.Digit3_png,
		&image_assets.Digit4_png,
		&image_assets.Digit5_png,
		&image_assets.Digit6_png,
		&image_assets.Digit7_png,
		&image_assets.Digit8_png,
		&image_assets.Digit9_png,
	} {
		img, _, err := image.Decode(bytes.NewReader(*raw))

		if err != nil {
			log.Fatal(err)
		}

		ui.digits[i] = ebiten.NewImageFromImage(img)
	}

	// Load Fonts
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)

	if err != nil {
		log.Fatal(err)
	}

	ui.text = etxt.NewStdRenderer()
	ui.text.SetFont(tt)
	ui.text.SetColor(color.Black)

}

func (ui *UI) DrawScore(screen *ebiten.Image, score uint) {

	const (
		XYOffset      = 8
		LetterSpacing = 2
	)

	op := ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(shared.GameWidth-XYOffset+LetterSpacing, XYOffset)

	for {
		ones := score % 10 // The ones' Place
		score /= 10        // Remove the ones place

		digit := ui.digits[ones]
		width := digit.Bounds().Dx()

		op.GeoM.Translate(-float64(width+LetterSpacing), 0)

		screen.DrawImage(digit, &op)

		if score == 0 {
			break
		}
	}
}

func (ui *UI) DrawHomeScreen(screen *ebiten.Image) {
	ui.text.SetTarget(screen)
	ui.text.SetAlign(etxt.YCenter, etxt.XCenter)
	ui.text.SetColor(color.RGBA{0xFF, 0x0, 0x0, 0xFF}) // RED

	ui.text.Draw(
		"Press Space\nto Start",
		shared.GameWidth/2,
		shared.GameHeight/4,
	)
}
