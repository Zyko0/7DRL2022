package graphics

import "github.com/hajimehoshi/ebiten/v2"

var (
	boxIndices = [6]uint16{0, 1, 2, 1, 2, 3}
)

func appendQuadVerticesIndices(vertices []ebiten.Vertex, indices []uint16, x, y, w, h, r, g, b, a float32, index int) ([]ebiten.Vertex, []uint16) {
	vertices = append(vertices, []ebiten.Vertex{
		{
			DstX:   x,
			DstY:   y,
			SrcX:   0,
			SrcY:   0,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x + w,
			DstY:   y,
			SrcX:   1,
			SrcY:   0,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x,
			DstY:   y + h,
			SrcX:   0,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
		{
			DstX:   x + w,
			DstY:   y + h,
			SrcX:   1,
			SrcY:   1,
			ColorR: r,
			ColorG: g,
			ColorB: b,
			ColorA: a,
		},
	}...)

	indiceCursor := uint16(index * 4)
	indices = append(indices, []uint16{
		boxIndices[0] + indiceCursor,
		boxIndices[1] + indiceCursor,
		boxIndices[2] + indiceCursor,
		boxIndices[3] + indiceCursor,
		boxIndices[4] + indiceCursor,
		boxIndices[5] + indiceCursor,
	}...)

	return vertices, indices
}
