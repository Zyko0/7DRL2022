package core

import (
	"math"

	"github.com/Zyko0/7DRL2022/core/entity"
	"github.com/Zyko0/7DRL2022/core/platform"
	"github.com/Zyko0/7DRL2022/logic"
)

func (c *Core) spawnRandomEnemy() {
	const (
		w = float64(logic.UnitSize)
		h = float64(logic.UnitSize)
	)

	x := -w
	if c.Player.X < logic.ScreenWidth/2 {
		x = logic.ScreenWidth + w
	}
	y := c.Player.Y + (c.rng.Float64() * 8 * logic.UnitSize)

	pathing := entity.PathingFollow
	if c.rng.Intn(2) == 0 {
		pathing = entity.PathingAnticipateY
	}

	enemy := entity.NewEnemy(x, y, w, h, entity.EnemySpec{
		Pathing:       pathing,
		MaxDistance:   byte(c.rng.Intn(int(c.Stats.MaxEnemyDistance + 1))),
		MoveSpeed:     c.Stats.EnemyMoveSpeed,
		ContactDamage: entity.ContactDamage1HP,
	})

	c.Entities = append(c.Entities, enemy)
}

func (c *Core) spawnChest(p *platform.Platform) {
	x := p.X
	y := p.Y + logic.UnitSize

	c.Entities = append(c.Entities, entity.NewChest(x, y))
}

func (c *Core) handleEntities() {
	px, py := c.Player.X, c.Player.Y
	pw, ph := c.Player.Width, float64(logic.UnitSize*2)
	for i, e := range c.Entities {
		e.Update(px, py)
		// Check contact with circle
		x, y := e.GetPosition()
		w, h := e.GetSize()
		// Skip unnecessary checks
		if y-h/2 < py+ph/2 && y+h/2 > py-ph/2 {
			// Check and apply contact
			if distance := math.Sqrt((x-px)*(x-px) + (y-py)*(y-py)); distance < w/2+pw/2 {
				if contact := e.Contact(); contact != entity.ContactNone {
					switch contact {
					case entity.ContactDamageHalfHP:
						c.Player.AddHP(-0.5)
					case entity.ContactDamage1HP:
						c.Player.AddHP(-1)
					case entity.ContactDamage2HP:
						c.Player.AddHP(-2)
					case entity.ContactHeal1HP:
						c.Player.AddHP(1)
					case entity.ContactItem:
						c.ChestPickedUp = true
					}
				}
			}
		}
		// If entity is destroyed, remove it
		if e.Destroyed() {
			c.Entities = append(c.Entities[:i], c.Entities[i+1:]...)
		}
	}
}
