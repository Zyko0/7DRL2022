package core

import (
	"github.com/Zyko0/7DRL2022/core/bonus"
	"github.com/Zyko0/7DRL2022/logic"
)

type Stats struct {
	AirControl          bool
	SpawnHearts         bool
	PlatformCellCount   int
	EventPlatformsCount int
	JumpForce           float64
	WaveHealMod         uint64
}

func NewStats() *Stats {
	return &Stats{
		AirControl:          false,
		SpawnHearts:         false,
		PlatformCellCount:   4,
		EventPlatformsCount: 4,
		JumpForce:           2,
		WaveHealMod:         0,
	}
}

func (s *Stats) ApplyBonus(b bonus.Bonus) {
	switch b {
	case bonus.BonusAirControl:
		s.AirControl = true
	case bonus.BonusSpawnHearts:
		s.SpawnHearts = true
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
