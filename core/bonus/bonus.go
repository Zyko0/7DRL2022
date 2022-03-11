package bonus

import (
	"math/rand"
)

type Bonus byte

const (
	BonusNone Bonus = iota
	BonusWaveHeal
	BonusAirControl
	BonusCrossBoundaries
	BonusSpawnHearts
	// BonusDash // Not sure
	// BonusDoubleJump // Not sure

	BonusStrongerJump
	BonusWeakerJump
)

var (
	primaries = []Bonus{
		BonusWaveHeal,
		BonusWaveHeal,
		BonusWaveHeal,
		BonusWaveHeal,
		BonusWaveHeal,
		BonusAirControl,
		BonusSpawnHearts,
		BonusCrossBoundaries,
	}

	secondaries = []Bonus{
		BonusStrongerJump,
		BonusStrongerJump,
		BonusStrongerJump,
		BonusStrongerJump,
		BonusStrongerJump,
		BonusWeakerJump,
		BonusWeakerJump,
		BonusWeakerJump,
	}
)

type List struct {
	rng *rand.Rand

	primaries   []Bonus
	secondaries []Bonus
}

func NewList(rng *rand.Rand) *List {
	p := make([]Bonus, len(primaries))
	s := make([]Bonus, len(secondaries))
	copy(p, primaries)
	copy(s, secondaries)

	return &List{
		rng: rng,

		primaries:   p,
		secondaries: s,
	}
}

func (bl *List) Roll() (Bonus, Bonus) {
	bl.rng.Shuffle(len(bl.primaries), func(i, j int) {
		bl.primaries[i], bl.primaries[j] = bl.primaries[j], bl.primaries[i]
	})
	bl.rng.Shuffle(len(bl.secondaries), func(i, j int) {
		bl.secondaries[i], bl.secondaries[j] = bl.secondaries[j], bl.secondaries[i]
	})

	return bl.primaries[0], bl.secondaries[0]
}

func (bl *List) Consume(b Bonus) {
	if b == BonusNone {
		return
	}

	if b == bl.primaries[0] {
		bl.primaries = bl.primaries[1:]
	} else {
		bl.secondaries = bl.secondaries[1:]
	}

	switch b {
	case BonusWeakerJump:
		bl.secondaries = append(bl.secondaries, BonusStrongerJump)
	case BonusStrongerJump:
		bl.secondaries = append(bl.secondaries, BonusWeakerJump)
	}
}
