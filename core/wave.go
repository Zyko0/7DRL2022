package core

import (
	"math/rand"

	"github.com/Zyko0/7DRL2022/logic"
)

const (
	WaveInitialGrowingRate = 1.0
	WaveMaxGrowingRate     = 2.0 // TODO: (?)
	WaveBaseHeight         = logic.ScreenHeight / 5
	WaveDamageTick         = logic.TPS / 2
)

type Wave struct {
	ticks      uint64
	rng        *rand.Rand
	rate       float64
	nextHeight float64

	healing bool
	levels  []float64
	min     float64
	max     float64
}

func NewWave(rng *rand.Rand) *Wave {
	levels := make([]float64, ColumnsCount)
	for i := range levels {
		levels[i] = WaveBaseHeight + rng.Float64()*WaveInitialGrowingRate
	}

	return &Wave{
		ticks:      0,
		rng:        rng,
		rate:       WaveInitialGrowingRate,
		nextHeight: 0,

		healing: false,
		levels:  levels,
	}
}

func (w *Wave) increaseGrowingRate() {
	if w.rate < WaveMaxGrowingRate {
		w.rate += 0.05
	}
}

func (w *Wave) advanceLevels(maxAmount float64) {
	min, max := float64(0), float64(0)
	for i := range w.levels {
		w.levels[i] += w.rng.Float64() * maxAmount
		if i == 0 || w.levels[i] < min {
			min = w.levels[i]
		}
		if i == 0 || w.levels[i] > max {
			max = w.levels[i]
		}
	}
	w.min = min
	w.max = max
}

func (w *Wave) Update(p *Player, healingMod uint64) {
	w.ticks++

	if healingMod > 0 {
		w.healing = w.ticks%(healingMod*2) < healingMod
	}
	if w.ticks%WaveDamageTick == 0 && p.Y+p.Height/2 < w.levels[int(p.X/60)] {
		dmg := -0.5
		if w.healing {
			dmg = 0.5
		}
		p.AddHP(dmg)
		// TODO: Play tick sound
	}

	if p.Height > w.nextHeight {
		w.increaseGrowingRate()
		w.nextHeight += WaveIncreaseFrequencyHeightInterval
	}

	w.advanceLevels(w.rate)
}

func (w *Wave) GetLevels() []float64 {
	return w.levels
}

func (w *Wave) GetMinMaxLevels() (float64, float64) {
	return w.min, w.max
}

func (w *Wave) Healing() bool {
	return w.healing
}
