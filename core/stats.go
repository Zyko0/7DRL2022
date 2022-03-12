package core

import (
	"github.com/Zyko0/7DRL2022/core/bonus"
	"github.com/Zyko0/7DRL2022/core/entity"
	"github.com/Zyko0/7DRL2022/core/platform"
	"github.com/Zyko0/7DRL2022/logic"
)

type Stats struct {
	AirControl          bool
	SpawnHearts         bool
	PlatformCellCount   int
	EventPlatformsCount int
	JumpForce           float64
	WaveHealMod         uint64
	EnemyMoveSpeed      float64
	MaxEnemyDistance    byte
}

func NewStats() *Stats {
	return &Stats{
		AirControl:          false,
		SpawnHearts:         false,
		PlatformCellCount:   platform.BaseCellsCount,
		EventPlatformsCount: 4,
		JumpForce:           BaseJumpForce,
		WaveHealMod:         0,
		EnemyMoveSpeed:      entity.MinEnemyMoveSpeed,
		MaxEnemyDistance:    entity.MaxDistanceShort,
	}
}

func (s *Stats) ApplyBonus(b bonus.Bonus) {
	switch b {
	case bonus.BonusAirControl:
		s.AirControl = true
	/*case bonus.BonusSpawnHearts:
	s.SpawnHearts = true*/
	case bonus.BonusWaveHeal:
		s.WaveHealMod += logic.TPS
	case bonus.BonusStrongerJump:
		s.JumpForce += 0.5
	case bonus.BonusWeakerJump:
		s.JumpForce -= 0.5
	}
}

func (s *Stats) Clone() *Stats {
	cloned := &Stats{}
	*cloned = *s

	return cloned
}
