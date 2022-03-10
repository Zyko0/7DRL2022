package entity

import "github.com/Zyko0/7DRL2022/logic"

type Chest struct {
	x float64
	y float64
	w float64
	h float64

	hadContact bool
}

func NewChest(x, y float64) Entity {
	return &Chest{
		x: 0,
		y: 0,
		w: logic.UnitSize,
		h: logic.UnitSize,

		hadContact: false,
	}
}

func (c *Chest) Update() {

}

func (c *Chest) Contact() Contact {
	if c.hadContact {
		return ContactNone
	}

	c.hadContact = true
	return ContactItem
}

func (c *Chest) GetPosition() (float64, float64) {
	return c.x, c.y
}

func (c *Chest) GetGfxParams() (float32, float32, float32, float32) {
	return 0, 0, 0, 0
}
