package bonus

func (b Bonus) String() string {
	switch b {
	case BonusNone:
		return "None"
	case BonusWaveHeal:
		return "Wave Heal"
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
