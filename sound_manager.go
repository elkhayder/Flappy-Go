package main

import (
	"bytes"
	"log"

	"github.com/elkhayder/Flappy-Go/assets/sounds"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type SoundManager struct {
	audioContext audio.Context

	background *audio.Player

	fx struct {
		die, jump, start, point *audio.Player
	}
}

func (sm *SoundManager) Init() {
	sm.audioContext = *audio.NewContext(44100)

	bgRaw, err := mp3.DecodeWithoutResampling(bytes.NewReader(sounds.Background_mp3))

	infinite := audio.NewInfiniteLoop(bgRaw, bgRaw.Length())

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

func (sm *SoundManager) Update() {
	// TODO : Infinite Loop
}
