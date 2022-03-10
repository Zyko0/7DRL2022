package entity

type Contact byte

const (
	ContactNone Contact = iota
	ContactDamage1HP
	ContactDamage2HP
	ContactHeal1HP
	ContactItem
)

type Entity interface {
	Update()
	Contact() Contact
	GetPosition() (float64, float64)
	GetGfxParams() (float32, float32, float32, float32)
}
