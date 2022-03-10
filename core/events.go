package core

import "github.com/Zyko0/7DRL2022/core/event"

func (c *Core) handleEvents() {
	for e := c.eventManager.ConsumeEvent(); e != event.KindNone; {
		switch e {
		case event.KindAoeSpawn:
		case event.KindEnemySpawn:
		case event.KindChestSpawn:
		case event.KindWaveReset:
		case event.KindSpecialPlatforms:
		default:
			break
		}
	}
}
