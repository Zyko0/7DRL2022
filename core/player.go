package core

import (
	"github.com/Zyko0/7DRL2022/core/platform"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	DefaultMoveSpeed = 5
	BaseJumpForce    = 2
)

type Player struct {
	X, Y          float64
	Width, Height float64
	MoveSpeed     float64
	JumpForce     float64

	HP               uint
	GroundedPlatform *platform.Platform
	Orientation      float64
	IntentVector     []float64
	VelocityVector   []float64
}

func NewPlayer() *Player {
	return &Player{
		X:         logic.ScreenWidth / 2,
		Y:         0,
		Width:     logic.UnitSize,
		Height:    logic.UnitSize * 2,
		MoveSpeed: DefaultMoveSpeed,
		JumpForce: BaseJumpForce,

		HP:               3,
		GroundedPlatform: nil,
		Orientation:      0,
		IntentVector:     []float64{0, 0},
		VelocityVector:   []float64{0, 0},
	}
}

func (p *Player) Update() {
	p.IntentVector[0] = 0
	p.IntentVector[1] = 0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.IntentVector[0] = -1
		p.Orientation = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.IntentVector[0] = 1
		p.Orientation = 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.IntentVector[1] = 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		p.Orientation = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		// Y++ required in order to detach from the current platform
		p.Y--
		p.IntentVector[1] = -1
		p.GroundedPlatform = nil
	}
}
