package graphics

import (
	"github.com/Zyko0/7DRL2022/assets"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	screenQuadVertices []ebiten.Vertex
	screenQuadIndices  []uint16
)

func init() {
	screenQuadVertices, screenQuadIndices = appendQuadVerticesIndices(
		screenQuadVertices, screenQuadIndices,
		0, 0,
		logic.ScreenWidth, logic.ScreenHeight,
		1, 1, 1, 1, 0,
	)
}

func (r *Renderer) RenderBackground(screen *ebiten.Image) {
	screen.DrawTrianglesShader(
		screenQuadVertices,
		screenQuadIndices,
		assets.BackgroundShader,
		&ebiten.DrawTrianglesShaderOptions{
			Uniforms: map[string]interface{}{
				"CameraY": float32(r.cameraYOffset),
				"ScreenSize": []float32{
					logic.ScreenWidth, logic.ScreenHeight,
				},
			},
		},
	)
}
