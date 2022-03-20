package assets

import (
	"bytes"
	"log"
	"math/rand"
	"time"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	defaultSFXVolume       = 1.0
	defaultGameMusicVolume = 0.4
	defaultMainMenuVolume  = 1.0
)

var (
	ctx = audio.NewContext(44100)

	//go:embed sfx/jump.wav
	jumpSoundBytes  []byte
	jumpSoundPlayer *audio.Player
	//go:embed sfx/jump2.wav
	jump2SoundBytes  []byte
	jump2SoundPlayer *audio.Player
	//go:embed sfx/jump3.wav
	jump3SoundBytes  []byte
	jump3SoundPlayer *audio.Player
)

func init() {
	var err error

	rand.Seed(time.Now().UnixNano())

	reader, err := wav.Decode(ctx, bytes.NewReader(jumpSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	jumpSoundPlayer, err = ctx.NewPlayer(reader)
	if err != nil {
		log.Fatal(err)
	}
	jumpSoundPlayer.SetVolume(defaultSFXVolume)

	reader, err = wav.Decode(ctx, bytes.NewReader(jump2SoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	jump2SoundPlayer, err = ctx.NewPlayer(reader)
	if err != nil {
		log.Fatal(err)
	}
	jump2SoundPlayer.SetVolume(defaultSFXVolume)

	reader, err = wav.Decode(ctx, bytes.NewReader(jump3SoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	jump3SoundPlayer, err = ctx.NewPlayer(reader)
	if err != nil {
		log.Fatal(err)
	}
	jump3SoundPlayer.SetVolume(defaultSFXVolume)
}

func PlayJumpSound() {
	var player *audio.Player

	switch rand.Intn(3) {
	case 0:
		player = jumpSoundPlayer
	case 1:
		player = jump2SoundPlayer
	case 2:
		player = jump3SoundPlayer
	}
	player.Rewind()
	player.Play()
}
