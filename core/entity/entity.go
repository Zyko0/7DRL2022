package entity

const (
	EntityChest = 0.
	EntityEnemy = 1.
	EntityAoe   = 2.
	EntityHeart = 3.
)

type Contact byte

const (
	ContactNone Contact = iota
	ContactDamageHalfHP
	ContactDamage1HP
	ContactDamage2HP
	ContactHeal1HP
	ContactItem
)

type Entity interface {
	Update(px, py float64)
	Contact() Contact
	Destroyed() bool
	GetPosition() (float64, float64)
	GetSize() (float64, float64)
	GetGfxParams() (float32, float32, float32, float32)
}

type base struct {
	x         float64
	y         float64
	w         float64
	h         float64
	destroyed bool
}

func (b *base) GetPosition() (float64, float64) {
	return b.x, b.y
}

func (b *base) GetSize() (float64, float64) {
	return b.w, b.h
}

func (b *base) Destroyed() bool {
	return b.destroyed
}
