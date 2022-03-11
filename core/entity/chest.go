package entity

import (
	"github.com/Zyko0/7DRL2022/logic"
)

type Chest struct {
	base

	hadContact bool
}

func NewChest(x, y float64) Entity {
	return &Chest{
		base: base{
			x:         x,
			y:         y,
			w:         logic.UnitSize,
			h:         logic.UnitSize,
			destroyed: false,
		},

		hadContact: false,
	}
}

func (c *Chest) Update(px, py float64) {}

func (c *Chest) Contact() Contact {
	if c.hadContact {
		return ContactNone
	}

	c.hadContact = true
	c.destroyed = true
	return ContactItem
}

func (c *Chest) GetGfxParams() (float32, float32, float32, float32) {
	return EntityChest, 0, 0, 0
}
