package utils

import (
	"math"

	"github.com/Zyko0/7DRL2022/logic"
)

const (
	FallSpeed    = -1.
	MaxFallSpeed = -logic.UnitSize + 1

	MinJumpVelocity  = 0.5
	MaxJumpVelocity  = 4.5
	baseJumpVelocity = 3.
)

func JumpVector(force, intentX float64) []float64 {
	v := []float64{intentX, baseJumpVelocity + force}
	length := math.Sqrt(v[0]*v[0] + v[1]*v[1])
	v[0] = v[0] / length * (baseJumpVelocity + force)
	v[1] = v[1] / length * (baseJumpVelocity + force)

	return v
}
