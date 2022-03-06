package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
	cameraYOffset float64
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Update(playerY float64) {
	r.cameraYOffset = playerY
}

func (r *Renderer) RenderBackground(screen *ebiten.Image) {
}
