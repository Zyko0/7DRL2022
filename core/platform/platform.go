package platform

import (
	"github.com/Zyko0/7DRL2022/logic"
)

const (
	MinCellsCount  = 2
	BaseCellsCount = 4
)

type Platform struct {
	sided float64

	GroundY  float64
	GroundNY float64

	X, Y       float64
	CellsCount int
	Width      float64
}

func New(x, y float64, cellsCount int, sided float64) *Platform {
	return &Platform{
		sided: sided,

		X:          x,
		Y:          y,
		CellsCount: cellsCount,
		Width:      float64(cellsCount) * logic.UnitSize,
	}
}
