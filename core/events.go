package core

import (
	"github.com/Zyko0/7DRL2022/core/entity"
	"github.com/Zyko0/7DRL2022/core/event"
	"github.com/Zyko0/7DRL2022/core/platform"
	"github.com/Zyko0/7DRL2022/core/utils"
	"github.com/Zyko0/7DRL2022/logic"
)

func (c *Core) handleEvents() {
	// TODO: remove below
	if c.Player.GroundedPlatform != nil {
		// c.spawnChest(c.Player.GroundedPlatform)
	}
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
				count := c.Stats.PlatformCellCount + c.rng.Intn(platform.BaseCellsCount-c.Stats.PlatformCellCount+1)
				p = platform.New(p.X, p.Y+logic.UnitSize*4, count, 0., true)
				c.Platforms.AddPlatform(p)
				missingRange = p.Y - c.Player.Y - logic.ScreenHeight*1.5
			}
		case event.KindReduceMaxPlatformWidth:
			c.Stats.PlatformCellCount--
			if c.Stats.PlatformCellCount < platform.MinCellsCount {
				c.Stats.PlatformCellCount = platform.MinCellsCount
			}
		case event.KindEnemyUpgrade:
			if c.Stats.MaxEnemyDistance < entity.MaxDistanceVeryLong {
				c.Stats.MaxEnemyDistance++
			}
		case event.KindEnemyFaster:
			if c.Stats.EnemyMoveSpeed < entity.MaxEnemyMoveSpeed {
				c.Stats.EnemyMoveSpeed += 0.5
			}
		}
	}
}
