package bonus

func (b Bonus) String() string {
	switch b {
	case BonusNone:
		return "None"
	case BonusWaveHeal:
		return "Wave Heal"
	case BonusWaveHeal2:
		return "Wave Heal 2"
	case BonusWaveHeal3:
		return "Wave Heal 3"
	case BonusWaveHeal4:
		return "Wave Heal 4"
	case BonusWaveHeal5:
		return "Wave Heal 5"
	case BonusAirControl:
		return "Air Control"
	case BonusCrossBoundaries:
		return "Cross Boundaries"
	case BonusSpawnHearts:
		return "Spawn Hearts"
	case BonusStrongerJump:
		return "Stronger Jump"
	case BonusWeakerJump:
		return "Weaker Jump"
	}

	return ""
}

func (b Bonus) Description() string {
	switch b {
	case BonusNone:
		return "No bonus"
	case BonusWaveHeal:
		return "The wave heals half the time"
	case BonusWaveHeal2:
		return "The wave heals half the time but longer"
	case BonusWaveHeal3:
		return "The wave heals half the time but longer"
	case BonusWaveHeal4:
		return "The wave heals half the time but longer"
	case BonusWaveHeal5:
		return "The wave heals half the time but longer"
	case BonusAirControl:
		return "Lateral movement control mid-air"
	case BonusCrossBoundaries:
		return "Cross the screen boundaries"
	case BonusSpawnHearts:
		return "Healing hearts can spawn"
	case BonusStrongerJump:
		return "Jump higher + trigger checkpoint"
	case BonusWeakerJump:
		return "Jump lower + trigger checkpoint"
	}

	return ""
}
