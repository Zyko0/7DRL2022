package graphics

import (
	"github.com/Zyko0/7DRL2022/assets"
	"github.com/Zyko0/7DRL2022/core/platform"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

func (r *Renderer) RenderPlatforms(screen *ebiten.Image, list *platform.List) {
	var (
		vertices []ebiten.Vertex
		indices  []uint16
	)

	y0 := r.cameraYOffset - logic.ScreenHeight/2 - logic.UnitSize
	y1 := y0 + logic.ScreenHeight + logic.UnitSize*2
	for i, p := range list.AppendPlatformsInRange(nil, y0, y1) {
		vertices, indices = appendQuadVerticesIndices(
			vertices,
			indices,
			float32(p.X-p.Width/2), float32(r.cameraYOffset-p.Y+logic.ScreenHeight/2-logic.UnitSize/2),
			float32(p.Width), logic.UnitSize,
			float32(p.CellsCount), 0, 0, 0,
			i,
		)
	}

	screen.DrawTrianglesShader(
		vertices,
		indices,
		assets.PlatformShader,
		&ebiten.DrawTrianglesShaderOptions{},
	)
}
