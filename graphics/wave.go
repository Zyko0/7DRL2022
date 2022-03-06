package graphics

import (
	"github.com/Zyko0/7DRL2022/assets"
	"github.com/Zyko0/7DRL2022/core"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

func (r *Renderer) RenderWave(screen *ebiten.Image, wave *core.Wave) {
	vertices, indices := appendQuadVerticesIndices(
		nil,
		nil,
		0, 0,
		logic.ScreenWidth, logic.ScreenHeight,
		0, 0, 0, 0,
		0,
	)

	levels := wave.GetLevels()
	uniforms := make([]float32, len(levels))
	for i := range levels {
		uniforms[i] = float32(levels[i] - r.cameraYOffset + logic.ScreenHeight/2)
	}

	min, max := wave.GetMinMaxLevels()
	screen.DrawTrianglesShader(vertices, indices, assets.WaveShader, &ebiten.DrawTrianglesShaderOptions{
		Uniforms: map[string]interface{}{
			"Levels":   uniforms,
			"MinLevel": float32(min - r.cameraYOffset + logic.ScreenHeight/2),
			"MaxLevel": float32(max - r.cameraYOffset + logic.ScreenHeight/2),
		},
	})
}
