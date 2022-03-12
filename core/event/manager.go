package event

import (
	"github.com/Zyko0/7DRL2022/logic"
)

type Manager struct {
	ticks  uint64
	height float64
	events []Kind

	lastAoeTick   uint64
	lastEnemyTick uint64

	lastChestHeight             uint64
	lastPlatformReductionHeight uint64
	lastEnemyUpgradeHeight      uint64
	lastEnemyFasterHeight       uint64
}

func NewManager() *Manager {
	return &Manager{
		ticks:  0,
		height: 0,
		events: make([]Kind, 0),

		lastAoeTick:   0,
		lastEnemyTick: 0,

		lastChestHeight: 0,
	}
}

func (m *Manager) Update(height float64) {
	m.ticks++
	if height > m.height {
		m.height = height
	}

	// Aoe
	const minAoeSpawnTick = 120 * logic.TPS
	if m.ticks >= minAoeSpawnTick {
		freq := tickAoeFrequencyFromHeight(m.height - minAoeSpawnTick)
		if m.ticks >= m.lastAoeTick+freq {
			m.lastAoeTick = m.ticks
			m.events = append(m.events, KindAoeSpawn)
		}
	}

	// Enemies
	const minEnemySpawnTick = 10 * logic.TPS
	if m.ticks >= minEnemySpawnTick {
		freq := tickEnemyFrequencyFromTicks(m.ticks)
		if m.ticks >= m.lastEnemyTick+freq {
			m.lastEnemyTick = m.ticks
			m.events = append(m.events, KindEnemySpawn)
		}
	}

	// Chest
	const minChestHeight = 6500
	if h := uint64(m.height - minChestHeight); h > 0 {
		if uint64(m.height) >= m.lastChestHeight+heightChestFrequency {
			m.lastChestHeight = uint64(m.height)
			m.events = append(m.events, KindChestSpawn)
			// Here we want to abort as there will be specific platforms above chest
			return
		}
	}

	// Reduce Platform width
	const minReducePlatformWidthHeight = 10000
	if h := uint64(m.height - minReducePlatformWidthHeight); h > 0 {
		if uint64(m.height) >= m.lastPlatformReductionHeight+heightPlatformReductionFrequency {
			m.lastPlatformReductionHeight = uint64(m.height)
			m.events = append(m.events, KindReduceMaxPlatformWidth)
		}
	}

	// Enemy Upgrade
	const minEnemyUpgradeHeight = 15000
	if h := uint64(m.height - minEnemyUpgradeHeight); h > 0 {
		if uint64(m.height) >= m.lastPlatformReductionHeight+heightEnemyUpgradeFrequency {
			m.lastEnemyUpgradeHeight = uint64(m.height)
			m.events = append(m.events, KindEnemyUpgrade)
		}
	}

	// Enemy Faster
	const minEnemyFasterHeight = 20000
	if h := uint64(m.height - minEnemyFasterHeight); h > 0 {
		if uint64(m.height) >= m.lastPlatformReductionHeight+heightEnemyFasterFrequency {
			m.lastEnemyFasterHeight = uint64(m.height)
			m.events = append(m.events, KindEnemyFaster)
		}
	}
}

func (m *Manager) ConsumeEvent() Kind {
	if len(m.events) == 0 {
		return KindNone
	}

	e := m.events[len(m.events)-1]
	m.events = m.events[0 : len(m.events)-1]

	return e
}
