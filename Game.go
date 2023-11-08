package main

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"strconv"
	"time"

	"github.com/elkhayder/Flappy-Go/assets/fonts"
	BackgroundSprites "github.com/elkhayder/Flappy-Go/assets/sprites/background"
	"github.com/elkhayder/Flappy-Go/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	scoreFont      font.Face
	mainScreenFont font.Face
)

type Game struct {
	Bird       entities.Bird
	Background *ebiten.Image
	pipes      [2]entities.Pipe

	score, highScore int
	ended            bool

	lastUpdatedAt int64
}

func (g *Game) Init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)

	if err != nil {
		log.Fatal(err)
	}

	scoreFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	mainScreenFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	bgImg, _, err := image.Decode(bytes.NewReader(BackgroundSprites.Day_png))

	if err != nil {
		log.Fatal(err)
	}

	// Load Background
	g.Background = ebiten.NewImageFromImage(bgImg)

	for i := range g.pipes {
		g.pipes[i] = entities.NewPipe(0, 0)
	}

	g.ended = true

}

func (g *Game) Reset() {
	width, height := g.Layout(0, 0)

	// Initiate the bird in middle of the screen
	g.Bird.Init(float64(width)/3, float64(height)/2)

	g.lastUpdatedAt = time.Now().UnixMicro()

	for i := range g.pipes {
		g.pipes[i].RandomizePositionOutsideRightScreen(width*(i+1), height)
	}

	g.ended = false

	if g.score > g.highScore {
		g.highScore = g.score
	}

	g.score = 0
}

func (g *Game) Update() error {
	if g.ended {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.Reset()
		} else {
			return nil
		}
	}

	ScreenWidth, ScreenHeight := g.Layout(0, 0)

	// Calculate dt
	now := time.Now().UnixMicro()
	dt := float64(now-g.lastUpdatedAt) * 1e-3 // Ms
	g.lastUpdatedAt = now

	// Update Bird
	g.Bird.Update(dt)
	if !g.Bird.CheckInsideScreen(ScreenWidth, ScreenHeight) {
		g.ended = true
		goto end
	}

	// Pipes
	for i := range g.pipes {
		pipe := &g.pipes[i]

		// Update
		pipe.Update(dt)

		// Check Collision and Halt the game
		if pipe.CheckCollision(&g.Bird) {
			g.ended = true
			goto end
		} else if pipe.IsBehindTheBird(&g.Bird) && !pipe.ScoreCounted {
			g.score++
			pipe.ScoreCounted = true
		} else if pipe.IsOutsideScreenLeft() {
			pipe.RandomizePositionOutsideRightScreen(ScreenWidth*2, ScreenHeight)
			pipe.ScoreCounted = false
		}
	}

end:
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ScreenWidth, ScreenHeight := g.Layout(0, 0)

	screen.DrawImage(g.Background, nil)

	if g.ended {

		DrawCenteredText(
			screen,
			"Press Space",
			mainScreenFont,
			ScreenWidth/2,
			ScreenHeight/3,
			color.RGBA{R: 0xFF, A: 0xFF},
		)
		DrawCenteredText(
			screen,
			" to Start!",
			mainScreenFont,
			ScreenWidth/2,
			ScreenHeight/3+30,
			color.RGBA{R: 0xFF, A: 0xFF},
		)
		DrawCenteredText(
			screen,
			"High Score: "+strconv.FormatInt(int64(g.highScore), 10),
			scoreFont,
			ScreenWidth/2,
			ScreenHeight/3+100,
			color.Black,
		)

		return
	}

	for _, p := range g.pipes {
		p.Draw(screen, nil)
	}

	g.Bird.Draw(screen)

	text.Draw(screen, "Score: "+strconv.FormatInt(int64(g.score), 10), scoreFont, 15, 30, color.Black)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 288, 512
}

func DrawCenteredText(screen *ebiten.Image, s string, font font.Face, cx, cy int, color color.Color) {
	bounds := text.BoundString(font, s)
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2
	text.Draw(screen, s, font, x, y, color)
}
