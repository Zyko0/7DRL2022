package assets

import (
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed shaders/wave.kage
	waveShaderSrc []byte
	//go:embed shaders/platform.kage
	platformShaderSrc []byte
	//go:embed shaders/sdentity.kage
	sdentityShaderSrc []byte
	//go:embed shaders/player.kage
	playerShaderSrc []byte
	//go:embed shaders/background.kage
	backgroundShaderSrc []byte

	WaveShader       *ebiten.Shader
	PlatformShader   *ebiten.Shader
	SDEntityShader   *ebiten.Shader
	PlayerShader     *ebiten.Shader
	BackgroundShader *ebiten.Shader
)

func init() {
	var err error

	WaveShader, err = ebiten.NewShader(waveShaderSrc)
	if err != nil {
		log.Fatal(err)
	}

	PlatformShader, err = ebiten.NewShader(platformShaderSrc)
	if err != nil {
		log.Fatal(err)
	}

	SDEntityShader, err = ebiten.NewShader(sdentityShaderSrc)
	if err != nil {
		log.Fatal(err)
	}

	PlayerShader, err = ebiten.NewShader(playerShaderSrc)
	if err != nil {
		log.Fatal(err)
	}

	BackgroundShader, err = ebiten.NewShader(backgroundShaderSrc)
	if err != nil {
		log.Fatal(err)
	}
}
