package graphics

import (
	"github.com/Zyko0/7DRL2022/assets"
	"github.com/Zyko0/7DRL2022/core"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

func (r *Renderer) RenderPlayer(screen *ebiten.Image, p *core.Player) {
	vertices, indices := appendQuadVerticesIndices(
		nil, nil,
		float32(p.X-p.Width/2), float32(logic.ScreenHeight/2-p.Height/2),
		float32(p.Width), float32(p.Height),
		0, 0, 0, 0, 0,
	)
	screen.DrawTrianglesShader(vertices, indices, assets.PlayerShader, &ebiten.DrawTrianglesShaderOptions{})
}
