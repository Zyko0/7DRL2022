package platform

import (
	"math/rand"

	"github.com/Zyko0/7DRL2022/core/utils"
	"github.com/Zyko0/7DRL2022/logic"
)

func Generate(rng *rand.Rand, from *Platform, cellsCount int, v []float64) *Platform {
	vx, vy := v[0], v[1]

	sided := 0.
	x0 := from.X - from.Width*0.45
	x1 := from.X + from.Width*0.45
	startX := rng.Float64()*(x1-x0) + x0
	startY := from.Y
	width := float64(cellsCount * logic.UnitSize)
	sign := 1.
	// Prevent a stack of platform on a side because it's boring
	if from.sided != 0. {
		sign = -from.sided
	} else if rng.Intn(2) == 0 {
		sign = -1.
	}
	vx *= sign
	// Find required iteration before fall
	sx := startX
	sy := startY
	minIteration := 0
	for vy > 0 {
		sx += vx
		sy += vy
		vy += utils.FallSpeed
		minIteration++
	}
	// Find max iterations to be above previous platform
	maxIteration := minIteration
	for sy > from.Y {
		sx += vx
		sy += vy
		vy += utils.FallSpeed
		maxIteration++
	}
	// Pick an iteration
	// TODO: -8 to avoid too much proximity => Need some tuning maybe
	iteration := minIteration + (rng.Intn(maxIteration - minIteration - 8))
	vy = v[1]
	for i := 0; i <= iteration; i++ {
		startX += vx
		startY += vy
		vy += utils.FallSpeed
		if vy < utils.MaxFallSpeed {
			vy = utils.MaxFallSpeed
		}
	}
	// Put it further or closer
	switch rng.Intn(2) {
	case 0:
		startX += rng.Float64() * sign * (width / 2)
	case 1:
		startX -= rng.Float64() * sign * (width / 2)
	}
	// Adjust to keep it on screen
	if startX < width/2 {
		startX = width / 2
		sided = sign
	} else if startX > logic.ScreenWidth-width/2 {
		startX = logic.ScreenWidth - width/2
		sided = sign
	}

	return new(startX, startY, cellsCount, sided)
}
