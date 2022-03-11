package ui

const (
	ActionNone = iota
	ActionResume
	ActionNewGame
)

type Main struct {
	action int
}

func NewMainMenu() *Main {
	return &Main{}
}

func (m *Main) Update() {
	// If right input, start a game or do crap
}

func (m *Main) GetAction() int {
	return ActionNewGame
}
