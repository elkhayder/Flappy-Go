package widgets

import (
	"fmt"
	"image"
	"log"

	gui_sprites "github.com/elkhayder/Flappy-Go/gui/sprites"
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

	for i := 0; i < 10; i++ {
		score.digits[i] = gui_sprites.Get(fmt.Sprintf("digits/%d.png", i))
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
