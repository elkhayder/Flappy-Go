package shared

import "github.com/hajimehoshi/ebiten/v2"

const (
	GameWidth  = 288
	GameHeight = 512
	TPS        = ebiten.DefaultTPS
	Dt         = 1.0 / TPS // dT
)
