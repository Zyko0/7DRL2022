package event

import "github.com/Zyko0/7DRL2022/logic"

const (
	tickWaveResetFrequency           = uint64(15 * logic.TPS) // TODO: not sure we want wave to reset
	heightChestFrequency             = uint64(6000)
	heightSpecialPlatformsFrequency  = uint64(5000)
	heightPlatformReductionFrequency = uint64(10000)
)

func tickEnemyFrequencyFromTicks(ticks uint64) uint64 {
	const (
		maxFreq   = logic.TPS
		minFreq   = 10 * logic.TPS
		incrEvery = 15 * logic.TPS
	)

	freq := minFreq - int64(ticks)*logic.TPS/incrEvery
	if freq < maxFreq {
		return maxFreq
	}

	return uint64(freq)
}

func tickAoeFrequencyFromHeight(height float64) uint64 {
	const (
		maxFreq   = logic.TPS
		minFreq   = logic.TPS * 10
		incrEvery = 5000
	)

	h := height - logic.ScreenHeight/2
	if h <= 0 {
		return minFreq
	}

	freq := uint64(minFreq - h/incrEvery*logic.TPS)
	if freq < maxFreq {
		return maxFreq
	}

	return freq
}
