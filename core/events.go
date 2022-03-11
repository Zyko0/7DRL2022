package core

import (
	"github.com/Zyko0/7DRL2022/core/event"
	"github.com/Zyko0/7DRL2022/core/platform"
	"github.com/Zyko0/7DRL2022/core/utils"
	"github.com/Zyko0/7DRL2022/logic"
)

func (c *Core) handleEvents() {
	for e := c.eventManager.ConsumeEvent(); e != event.KindNone; e = c.eventManager.ConsumeEvent() {
		switch e {
		case event.KindAoeSpawn:
		case event.KindEnemySpawn:
			c.spawnRandomEnemy()
		case event.KindChestSpawn:
			p := c.Platforms.GetHighestPlatform()
			c.lastChestPlatform = p
			c.spawnChest(p)
			// Append friendly platforms after a potential chess lowering jump force
			missingRange := p.Y - c.Player.Y - logic.ScreenHeight*1.5
			jumpv := utils.JumpVector(c.Stats.JumpForce, 1)
			jumpv[0] *= c.Player.MoveSpeed
			jumpv[1] *= c.Player.MoveSpeed
			// Generate missing platforms in advance
			for missingRange < 0 {
				p = platform.New(p.X, p.Y+logic.UnitSize*4, c.Stats.PlatformCellCount, 0.)
				c.Platforms.AddPlatform(p)
				missingRange = p.Y - c.Player.Y - logic.ScreenHeight*1.5
			}
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
