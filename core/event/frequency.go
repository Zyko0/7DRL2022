package event

import "github.com/Zyko0/7DRL2022/logic"

const (
	tickWaveResetFrequency          = uint64(15 * logic.TPS)
	heightChestFrequency            = uint64(6500)
	heightSpecialPlatformsFrequency = uint64(5000)
)

func tickEnemyFrequencyFromHeight(height float64) uint64 {
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
	if freq > maxFreq {
		return maxFreq
	}

	return freq
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
	if freq > maxFreq {
		return maxFreq
	}

	return freq
}
