package graphics

import (
	"github.com/Zyko0/7DRL2022/assets"
	"github.com/Zyko0/7DRL2022/core/entity"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/hajimehoshi/ebiten/v2"
)

func (r *Renderer) RenderEntities(screen *ebiten.Image, entities []entity.Entity) {
	var (
		vertices []ebiten.Vertex
		indices  []uint16
	)

	y0 := r.cameraYOffset - logic.ScreenHeight/2 - logic.UnitSize
	y1 := y0 + logic.ScreenHeight + logic.UnitSize*2
	index := 0
	for _, e := range entities {
		x, y := e.GetPosition()
		if y < y0 || y > y1 {
			continue
		}
		w, h := e.GetSize()
		pr, pg, pb, pa := e.GetGfxParams()
		vertices, indices = AppendQuadVerticesIndices(
			vertices,
			indices,
			float32(x-w/2), float32(r.cameraYOffset-y+logic.ScreenHeight/2-h/2),
			float32(w), float32(h),
			pr, pg, pb, pa,
			index,
		)
		index++
	}

	screen.DrawTrianglesShader(
		vertices,
		indices,
		assets.SDEntityShader,
		&ebiten.DrawTrianglesShaderOptions{},
	)
}
