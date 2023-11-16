package widgets

import (
	"image"

	gui_sprites "github.com/elkhayder/Flappy-Go/gui/sprites"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/furex/v2"
)

type Slider struct {
	target  *float64
	clicked bool

	frame image.Rectangle
}

var (
	_ furex.Updater       = (*Slider)(nil)
	_ furex.Drawer        = (*Slider)(nil)
	_ furex.MouseHandler  = (*Slider)(nil)
	_ furex.ButtonHandler = (*Slider)(nil)
)

func NewSlider(target *float64) *Slider {
	return &Slider{
		target: target,
	}
}

func (s *Slider) Update(v *furex.View) {
	//TODO

}

func (s *Slider) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	s.frame = frame // Update the frame, used to implement sliding functionnality

	var (
		StartSprite, EndSprite *ebiten.Image

		FilledMid = gui_sprites.Get("Slider_Filled_Mid.png")
		EmptyMid  = gui_sprites.Get("Slider_Empty_Mid.png")
		Handle    = gui_sprites.Get("Slider_Handle.png")
	)

	// Load Start Sprite
	if *s.target == 0 {
		StartSprite = gui_sprites.Get("Slider_Empty_Start.png")
	} else {
		StartSprite = gui_sprites.Get("Slider_Filled_Start.png")
	}

	// Load End Sprite
	if *s.target == 1 {
		EndSprite = gui_sprites.Get("Slider_Filled_End.png")
	} else {
		EndSprite = gui_sprites.Get("Slider_Empty_End.png")
	}

	width := float64(frame.Dx())
	drawnWidth := 0.0

	HeightScale := float64(frame.Dy()) / float64(StartSprite.Bounds().Dy())

	op := ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Scale(1, HeightScale)

	op.GeoM.Translate(float64(frame.Min.X), float64(frame.Min.Y))

	// Draw Start
	screen.DrawImage(StartSprite, &op)
	drawnWidth += float64(StartSprite.Bounds().Dx())
	op.GeoM.Translate(float64(StartSprite.Bounds().Dx()), 0)

	// Draw the middle
	for drawnWidth < width-float64(EndSprite.Bounds().Dx()) {
		var sprite *ebiten.Image

		if drawnWidth/width < *s.target {
			sprite = FilledMid
		} else {
			sprite = EmptyMid
		}

		screen.DrawImage(sprite, &op)
		op.GeoM.Translate(1, 0)
		drawnWidth += 1
	}

	// Draw the End
	screen.DrawImage(EndSprite, &op)

	// Draw Handle
	op.GeoM.Reset()

	x := float64(frame.Min.X) + (*s.target)*width - float64(Handle.Bounds().Dx())/2

	op.GeoM.Scale(1, float64(frame.Dy())/float64(Handle.Bounds().Dy()))
	op.GeoM.Translate(x, float64(frame.Min.Y))

	screen.DrawImage(Handle, &op)
}

func (s *Slider) HandlePress(x, y int, t ebiten.TouchID) {
	s.clicked = true
	s.HandleMouse(x, y)
}

func (s *Slider) HandleRelease(x, y int, isCancel bool) {
	s.clicked = false
}

func (s *Slider) HandleMouse(x, y int) bool {
	if !s.clicked {
		return false
	}

	if x < s.frame.Min.X {
		*s.target = 0
	} else if x > s.frame.Max.X {
		*s.target = 1
	} else {
		*s.target = float64(x-s.frame.Min.X) / float64(s.frame.Dx())
	}

	return true
}
