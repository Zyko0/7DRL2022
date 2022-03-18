package assets

import (
	"bytes"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	defaultSFXVolume       = 2.0
	defaultGameMusicVolume = 0.4
	defaultMainMenuVolume  = 1.0
)

var (
	ctx = audio.NewContext(44100)

	//go:embed sfx/jump.wav
	jumpSoundBytes  []byte
	jumpSoundPlayer *audio.Player
)

func init() {
	var err error

	reader, err := wav.Decode(ctx, bytes.NewReader(jumpSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	jumpSoundPlayer, err = ctx.NewPlayer(reader)
	if err != nil {
		log.Fatal(err)
	}
	jumpSoundPlayer.SetVolume(defaultSFXVolume)
}

func PlayJumpSound() {
	jumpSoundPlayer.Rewind()
	jumpSoundPlayer.Play()
}
