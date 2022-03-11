package event

import (
	"github.com/Zyko0/7DRL2022/logic"
)

type Manager struct {
	ticks  uint64
	height float64
	events []Kind

	lastAoeTick       uint64
	lastEnemyTick     uint64
	lastWaveResetTick uint64

	lastChestHeight            uint64
	lastSpecialPlatformsHeight uint64
}

func NewManager() *Manager {
	return &Manager{
		ticks:  0,
		height: 0,
		events: make([]Kind, 0),

		lastAoeTick:       0,
		lastEnemyTick:     0,
		lastWaveResetTick: 0,

		lastChestHeight:            0,
		lastSpecialPlatformsHeight: 0,
	}
}

func (m *Manager) Update(height float64) {
	m.ticks++
	if height > m.height {
		m.height = height
	}

	// Aoe
	const minAoeSpawnTick = 120 * logic.TPS
	if m.ticks > minAoeSpawnTick {
		freq := tickAoeFrequencyFromHeight(m.height - minAoeSpawnTick)
		if m.ticks >= m.lastAoeTick+freq {
			m.lastAoeTick = m.ticks
			m.events = append(m.events, KindAoeSpawn)
		}
	}

	// Enemies
	const minEnemySpawnTick = 60 * logic.TPS
	// TODO: if m.ticks > minEnemySpawnTick
	freq := tickEnemyFrequencyFromTicks(m.ticks)
	if m.ticks >= m.lastEnemyTick+freq {
		m.lastEnemyTick = m.ticks
		m.events = append(m.events, KindEnemySpawn)
	}

	// Wave reset
	freq = tickWaveResetFrequency
	if m.ticks >= m.lastWaveResetTick+freq {
		m.lastWaveResetTick = m.ticks
		m.events = append(m.events, KindWaveReset)
	}

	// Chest
	const minChestHeight = 8500
	if h := uint64(m.height - minChestHeight); h > 0 {
		if h >= m.lastChestHeight+heightChestFrequency {
			m.lastChestHeight = uint64(m.height)
			m.events = append(m.events, KindChestSpawn)
			// Here we want to abort as there will be specific platforms above chest
			return
		}
	}

	// Special platforms
	const minSpecialPlatformsHeight = 10000
	if h := uint64(m.height - minSpecialPlatformsHeight); h > 0 {
		if h >= m.lastSpecialPlatformsHeight+heightSpecialPlatformsFrequency {
			m.lastSpecialPlatformsHeight = uint64(m.height)
			m.events = append(m.events, KindSpecialPlatforms)
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
