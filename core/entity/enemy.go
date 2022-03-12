package entity

import (
	"math"

	"github.com/Zyko0/7DRL2022/logic"
)

const (
	MinEnemyMoveSpeed = 4
	MaxEnemyMoveSpeed = 8
)

const (
	PathingFollow byte = iota
	PathingAnticipateY
)

const (
	MaxDistanceShort byte = iota
	MaxDistanceMedium
	MaxDistanceLong
	MaxDistanceVeryLong
)

type EnemySpec struct {
	Pathing       byte
	MaxDistance   byte
	MoveSpeed     float64
	ContactDamage Contact
}

type Enemy struct {
	base

	ticks            uint64
	initialDeathTick int64
	deathTick        int64
	pathing          byte
	maxDist          byte
	moveSpeed        float64
	dmg              Contact

	hadContact bool
}

func NewEnemy(x, y, w, h float64, spec EnemySpec) Entity {
	return &Enemy{
		base: base{
			x:         x,
			y:         y,
			w:         w,
			h:         h,
			destroyed: false,
		},

		ticks:            0,
		maxDist:          spec.MaxDistance,
		initialDeathTick: 0,
		deathTick:        -1,
		pathing:          spec.Pathing,
		moveSpeed:        spec.MoveSpeed,
		dmg:              spec.ContactDamage,

		hadContact: false,
	}
}

func (e *Enemy) Update(px, py float64) {
	if e.deathTick == 0 || e.hadContact {
		e.destroyed = true
		e.deathTick = 0
		return
	}

	// Init death tick
	if e.initialDeathTick == 0 {
		dist := math.Sqrt((e.x-px)*(e.x-px) + (e.y-py)*(e.y-py))
		switch e.maxDist {
		case MaxDistanceShort:
			dist *= 1.25
		case MaxDistanceMedium:
			dist *= 1.5
		case MaxDistanceLong:
			dist *= 2.
		case MaxDistanceVeryLong:
			dist *= 3.
		}
		dist /= e.moveSpeed
		e.initialDeathTick = int64(dist)
		e.deathTick = e.initialDeathTick
	}

	// Apply velocity
	vx, vy := 0., 0.
	switch e.pathing {
	case PathingFollow:
		vx = px - e.x
		vy = py - e.y
	case PathingAnticipateY:
		vx = px - e.x
		vy = py + (logic.UnitSize * 4) - e.y
		// Edge case if already close to same X axis
		if math.Abs(vx) <= 10. {
			vy = py - e.y
		}
	}
	length := math.Sqrt(vx*vx + vy*vy)
	e.x = e.x + vx/length*e.moveSpeed
	e.y = e.y + vy/length*e.moveSpeed

	e.deathTick--
	e.ticks++
}

func (e *Enemy) Contact() Contact {
	if e.hadContact {
		return ContactNone
	}

	e.hadContact = true
	e.destroyed = true
	return ContactDamage1HP
}

func (e *Enemy) GetGfxParams() (float32, float32, float32, float32) {
	return EntityEnemy,
		float32(e.maxDist) / float32(MaxDistanceVeryLong),
		float32((e.moveSpeed - MinEnemyMoveSpeed) / (MaxEnemyMoveSpeed - MinEnemyMoveSpeed)),
		float32(e.deathTick) / float32(e.initialDeathTick)
}
