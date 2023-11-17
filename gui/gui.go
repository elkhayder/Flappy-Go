package gui

import (
	_ "embed"
	"log"

	"github.com/elkhayder/Flappy-Go/assets/fonts"
	"github.com/elkhayder/Flappy-Go/config"
	gui_sprites "github.com/elkhayder/Flappy-Go/gui/sprites"
	"github.com/elkhayder/Flappy-Go/gui/widgets"
	"github.com/elkhayder/Flappy-Go/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"golang.org/x/image/font/opentype"
)

type ViewType int

const (
	ViewTypeSettings ViewType = iota
	ViewTypeScore
)

type GUI struct {
	ui   *furex.View
	text *etxt.Renderer

	CurrentView ViewType
}

//go:embed settings.html
var settingsHTML string

var SettingsView, HomeView, ScoreView *furex.View

func (gui *GUI) Init() {
	gui_sprites.Load()

	// Load Fonts
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)

	if err != nil {
		log.Fatal(err)
	}

	gui.CurrentView = ViewTypeScore

	gui.text = etxt.NewStdRenderer()
	gui.text.SetFont(tt)

	SettingsView = furex.Parse(settingsHTML, &furex.ParseOptions{
		Width:  shared.GameWidth,
		Height: shared.GameHeight,
		Components: furex.ComponentsMap{
			"card": &widgets.Card{},
			"text": widgets.NewText(
				gui.text,
				nil,
				14,
				etxt.Left,
				etxt.YCenter,
			),
			"fx-enabled-checkbox": widgets.NewCheckbox(
				&config.Config.Fx.Enabled,
			),
			"music-enabled-checkbox": widgets.NewCheckbox(
				&config.Config.Music.Enabled,
			),
			"fx-volume-slider": widgets.NewSlider(
				&config.Config.Fx.Volume,
			),
			"music-volume-slider": widgets.NewSlider(
				&config.Config.Music.Volume,
			),
		},
	})

	ScoreView = &furex.View{
		Handler: widgets.NewScore(),
	}

}

func (gui *GUI) Update() {
	if gui.CurrentView == ViewTypeSettings && gui.ui != SettingsView {
		gui.ui = SettingsView
	} else if gui.CurrentView == ViewTypeScore && gui.ui != ScoreView {
		gui.ui = ScoreView
	}

	gui.ui.Update()
}

func (gui *GUI) Draw(screen *ebiten.Image) {
	gui.ui.Draw(screen)
}
