package widgets

import (
	"bytes"
	"image"
	"log"

	image_assets "github.com/elkhayder/Flappy-Go/assets/images"
	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
)

var ScorePointer *uint = nil

type Score struct {
	digits [10]*ebiten.Image
}

var (
	_ furex.Drawer = (*Score)(nil)
)

func NewScore() *Score {
	score := Score{}

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

		score.digits[i] = ebiten.NewImageFromImage(img)
	}

	return &score
}

func (s *Score) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	if ScorePointer == nil {
		log.Fatal("Score pointer is nil")
	}

	score := *ScorePointer

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

		digit := s.digits[ones]
		width := digit.Bounds().Dx()

		op.GeoM.Translate(-float64(width+LetterSpacing), 0)

		screen.DrawImage(digit, &op)

		if score == 0 {
			break
		}
	}
}
