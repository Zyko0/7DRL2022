package platform

import (
	"math/rand"

	"github.com/Zyko0/7DRL2022/logic"
)

type List struct {
	indices   [][]int
	platforms []*Platform
}

func NewList() *List {
	return &List{
		indices:   make([][]int, 0),
		platforms: make([]*Platform, 0),
	}
}

func (pl *List) Initialize(rng *rand.Rand) {
	// Add a base platform
	pl.AddPlatform(new(
		rng.Float64()*(logic.ScreenWidth-BaseCellsCount),
		logic.ScreenHeight/2+logic.UnitSize/2,
		BaseCellsCount,
		0,
	))
}

func (l *List) AddPlatform(p *Platform) {
	// Y + platformHeight/2 + playerHeight/2
	p.GroundY = p.Y + logic.UnitSize/2 + logic.UnitSize
	p.GroundNY = p.Y - logic.UnitSize/2 - logic.UnitSize

	l.platforms = append(l.platforms, p)

	index := int(p.Y) / 1000
	// Allocate in advance
	for index > len(l.indices)-1 {
		l.indices = append(l.indices, []int{})
	}
	l.indices[index] = append(l.indices[index], len(l.platforms)-1)
}

func (l *List) AppendPlatformsInRange(platforms []*Platform, y0, y1 float64) []*Platform {
	index0 := int(y0) / 1000
	indexn := int(y1) / 1000
	for index := index0; index <= indexn && index <= len(l.indices)-1; index++ {
		for _, pi := range l.indices[index] {
			py0 := l.platforms[pi].Y - logic.UnitSize/2
			py1 := l.platforms[pi].Y + logic.UnitSize/2
			if (py0 >= y0 && py0 <= y1) || (py1 >= y0 && py1 <= y1) {
				platforms = append(platforms, l.platforms[pi])
			}
		}
	}

	return platforms
}

func (l *List) GetHighestPlatform() *Platform {
	var platform *Platform

	maxY := 0.
	for _, pi := range l.indices[len(l.indices)-1] {
		if p := l.platforms[pi]; p.Y > maxY {
			platform = p
			maxY = p.Y
		}
	}

	return platform
}

func (l *List) Slice() []*Platform {
	return l.platforms
}

func (l *List) Count() int {
	return len(l.platforms)
}
