package core

import (
	"github.com/Zyko0/7DRL2022/core/event"
	"github.com/Zyko0/7DRL2022/core/platform"
)

func (c *Core) handleEvents() {
	for e := c.eventManager.ConsumeEvent(); e != event.KindNone; e = c.eventManager.ConsumeEvent() {
		switch e {
		case event.KindAoeSpawn:
		case event.KindEnemySpawn:
			c.spawnRandomEnemy()
		case event.KindChestSpawn:
			p := c.Platforms.GetHighestPlatform()
			c.spawnChest(p)
		case event.KindWaveReset:
		case event.KindSpecialPlatforms:
		case event.KindReducePlatformWidth:
			c.Stats.PlatformCellCount--
			if c.Stats.PlatformCellCount < platform.MinCellsCount {
				c.Stats.PlatformCellCount = platform.MinCellsCount
			}
		default:
		}
	}
}
