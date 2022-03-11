package entity

import (
	"github.com/Zyko0/7DRL2022/logic"
)

type Heart struct {
	base

	hadContact bool
}

func NewHeart(x, y float64) Entity {
	return &Heart{
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

func (h *Heart) Update(px, py float64) {
	if h.hadContact {
		h.destroyed = true
		return
	}
}

func (h *Heart) Contact() Contact {
	if h.hadContact {
		return ContactNone
	}

	h.hadContact = true
	h.destroyed = true
	return ContactHeal1HP
}

func (h *Heart) GetGfxParams() (float32, float32, float32, float32) {
	return EntityHeart, 0, 0, 0
}
