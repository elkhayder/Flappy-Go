package config

import (
	"bytes"
	"log"

	"github.com/elkhayder/Flappy-Go/assets/sounds"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type SoundManager struct {
	audioContext audio.Context

	background *audio.Player

	fx struct {
		die, jump, start, point *audio.Player
	}
}

type Fx uint

const (
	FxDie Fx = iota
	FxJump
	FxStart
	FxPoint
)

func (sm *SoundManager) Init() {
	sm.audioContext = *audio.NewContext(44100)

	bgRaw, err := wav.DecodeWithoutResampling(bytes.NewReader(sounds.ChumbucketRhumba_wav))
	if err != nil {
		log.Fatal(err)
	}

	const (
		BackgroundMusicDuration  = 13.618 //s
		BackgroundIntroDuration  = 1.8    // s
		BackgroundIntroLoopRatio = BackgroundIntroDuration / BackgroundMusicDuration
	)

	introSize := BackgroundIntroLoopRatio * float64(bgRaw.Length())

	infinite := audio.NewInfiniteLoopWithIntro(bgRaw, int64(introSize), bgRaw.Length())
	if err != nil {
		log.Fatal(err)
	}

	sm.background, err = sm.audioContext.NewPlayer(infinite)
	if err != nil {
		log.Fatal(err)
	}

	/// Load Effects

	fxs := []struct {
		raw         *[]byte
		destination **audio.Player
	}{
		// {raw: &sounds.Die_wav, destination: &sm.fx.hitPipe},
		{raw: &sounds.Hit_wav, destination: &sm.fx.die},
		{raw: &sounds.Wing_wav, destination: &sm.fx.jump},
		{raw: &sounds.Swoosh_wav, destination: &sm.fx.start},
		{raw: &sounds.Point_wav, destination: &sm.fx.point},
	}

	for i := range fxs {
		raw, err := wav.DecodeWithoutResampling(bytes.NewReader(*fxs[i].raw))

		if err != nil {
			log.Fatal(err)
		}

		*fxs[i].destination, err = sm.audioContext.NewPlayer(raw)

		if err != nil {
			log.Fatal(err)
		}
	}

}

func (sm *SoundManager) PlayFx(fx Fx) {
	if !Config.Fx.Enabled {
		return
	}

	var player *audio.Player

	switch fx {
	case FxDie:
		player = sm.fx.die
	case FxJump:
		player = sm.fx.jump
	case FxPoint:
		player = sm.fx.point
	case FxStart:
		player = sm.fx.start

	}

	player.Rewind()
	player.SetVolume(Config.Fx.Volume)

	player.Play()
}

func (sm *SoundManager) Update() {
	if Config.Music.Enabled && !sm.background.IsPlaying() {
		sm.background.Play()
	}

	if !Config.Music.Enabled && sm.background.IsPlaying() {
		sm.background.Pause()
	}

	if sm.background.Volume() != Config.Music.Volume {
		sm.background.SetVolume(Config.Music.Volume)
	}

	for _, p := range []*audio.Player{
		sm.fx.die,
		sm.fx.jump,
		sm.fx.start,
		sm.fx.point,
	} {
		if p.Volume() != Config.Fx.Volume {
			p.SetVolume(Config.Fx.Volume)
		}
	}
}
