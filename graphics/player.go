package graphics

import (
	"github.com/Zyko0/7DRL2022/assets"
	"github.com/Zyko0/7DRL2022/core"
	"github.com/Zyko0/7DRL2022/core/utils"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

func (r *Renderer) RenderPlayer(screen *ebiten.Image, p *core.Player) {
	vertices, indices := AppendQuadVerticesIndices(
		nil, nil,
		float32(p.X-p.Width/2),
		float32(logic.ScreenHeight/2-p.Height/2),
		float32(p.Width), float32(p.Height),
		0, 0, 0, 0, 0,
	)

	orientationX, orientationY := float32(p.Orientation), float32(0)
	// If falling at max speed, force looking down for fun
	if p.VelocityVector[1] <= utils.MaxFallSpeed {
		orientationX = 0
		// TODO: if air control available, don't ! since X orientation is a sensitive info
		orientationY = -1
	}
	screen.DrawTrianglesShader(vertices, indices, assets.PlayerShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: map[string]interface{}{
			"Orientation": []float32{orientationX, orientationY},
			"Ticks":       float32(p.JumpTicks),
		},
	})
}
