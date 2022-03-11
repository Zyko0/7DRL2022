package graphics

import (
	"github.com/Zyko0/7DRL2022/core"
	"github.com/hajimehoshi/ebiten/v2"
)

func (r *Renderer) drawHPBar(screen *ebiten.Image, hp float64) {
	const (
		OffsetX     = 50
		OffsetY     = 25
		BorderWidth = 2
		Width       = 300
		Height      = 50
	)

	// Container
	vertices, indices := AppendQuadVerticesIndices(
		nil, nil,
		float32(OffsetX), float32(OffsetY),
		float32(Width), float32(Height),
		0, 0, 0, 0.75, 0,
	)
	// Bar
	vertices, indices = AppendQuadVerticesIndices(
		vertices, indices,
		float32(OffsetX+BorderWidth), float32(OffsetY+BorderWidth),
		float32(Width-BorderWidth*2)*float32(hp/core.PlayerMaxHP), float32(Height-BorderWidth*2),
		0.75, 0.25, 0.5, 0.75, 1,
	)

	screen.DrawTriangles(vertices, indices, brushImage, nil)
}

func (r *Renderer) RenderHUD(screen *ebiten.Image, hp float64, score uint64) {
	r.drawHPBar(screen, hp)
}
