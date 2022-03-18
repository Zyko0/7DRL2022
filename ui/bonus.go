package ui

import (
	"github.com/Zyko0/7DRL2022/assets"
	"github.com/Zyko0/7DRL2022/core/bonus"
	"github.com/Zyko0/7DRL2022/graphics"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	bonusBgWidth   = 0.8 * logic.ScreenHeight
	bonusBgHeight  = 0.4 * logic.ScreenHeight
	bonusBgOffsetX = (logic.ScreenWidth - bonusBgWidth) / 2
	bonusBgOffsetY = (logic.ScreenHeight - bonusBgHeight) / 2

	bonusCardIntervalOffset = 12
	bonusCardWidth          = (bonusBgWidth - bonusCardIntervalOffset*3) / 2
	bonusCardHeight         = bonusBgHeight/6*4 - bonusCardIntervalOffset
	bonusCardOffsetY        = bonusBgHeight / 6

	skipCardWidth = bonusCardWidth
)

var (
	bonusBgColor = []float32{0.7, 0, 0, 0.7}
)

type BonusView struct {
	active bool

	lastCursorX int
	lastCursorY int
	card        *ebiten.Image

	SelectedIndex int
	Bonuses       []bonus.Bonus
}

func NewBonusView() *BonusView {
	card := ebiten.NewImage(bonusBgWidth, bonusBgHeight)

	graphics.DrawRect(card, 0, 0, bonusBgWidth, bonusBgHeight, 0.1, 0.1, 0.1, 0.9)
	graphics.DrawRectBorder(card, 0, 0, bonusBgWidth, bonusBgHeight, 1, 1, 1, 1, 0.9)
	// Title
	str := "Pick a bonus"
	rect := text.BoundString(assets.CardTitleFontFace, str)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(bonusBgWidth/2-rect.Max.X/2),
		float64(48),
	)
	text.DrawWithOptions(card, str, assets.CardTitleFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	return &BonusView{
		active: false,

		card: card,

		SelectedIndex: 0,
		Bonuses:       nil,
	}
}

func (bv *BonusView) Reset() {
	bv.active = false
	bv.SelectedIndex = 0
}

func (bv *BonusView) Active() bool {
	return bv.active
}

func (bv *BonusView) SetBonuses(bonuses []bonus.Bonus) {
	bv.Bonuses = bonuses
	bv.active = true
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
}

func (bv *BonusView) Update() {
	// Update selection based on keyboard input at last
	var kbInput bool

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		kbInput = true
		bv.SelectedIndex++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		kbInput = true
		bv.SelectedIndex--
	}
	if bv.SelectedIndex < 0 {
		bv.SelectedIndex = len(bv.Bonuses) - 1
	}
	if bv.SelectedIndex >= len(bv.Bonuses) {
		bv.SelectedIndex = 0
	}

	// Check what is on mouse hover only if there has not been any keyboard input and an actual mouse input
	var hovered bool

	x, y := ebiten.CursorPosition()
	hoveredIndex := 0
	y0 := bonusBgOffsetY + bonusCardOffsetY
	y1 := y0 + bonusCardHeight
	for i := 0; i < len(bv.Bonuses)-1; i++ {
		x0 := bonusBgOffsetX + float64(i+1)*bonusCardIntervalOffset + float64(i)*bonusCardWidth
		x1 := x0 + bonusCardWidth
		if float64(x) >= x0 && float64(x) <= x1 && float64(y) >= y0 && float64(y) <= y1 {
			hoveredIndex = i
			hovered = true
			break
		}
	}
	// Check for hover on skip button
	y0 = y1 + bonusCardIntervalOffset
	y1 = y0 + (bonusCardHeight/2 - bonusCardIntervalOffset)
	x0 := bonusBgOffsetX + bonusCardIntervalOffset
	x1 := x0 + bonusCardIntervalOffset + bonusCardWidth*2
	if float64(x) >= x0 && float64(x) <= x1 && float64(y) >= y0 && float64(y) <= y1 {
		hoveredIndex = len(bv.Bonuses) - 1
		hovered = true
	}

	if !kbInput && (x != bv.lastCursorX || y != bv.lastCursorY) {
		bv.SelectedIndex = hoveredIndex
	}

	// Check if a selection is made
	var validated bool

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		validated = true
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && hovered {
		validated = true
	}
	if validated {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		bv.active = false
	}

	bv.lastCursorX, bv.lastCursorY = x, y
}

func (bv *BonusView) Draw(screen *ebiten.Image) {
	// Background card
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(bonusBgOffsetX),
		float64(bonusBgOffsetY),
	)
	screen.DrawImage(bv.card, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Augment cards
	for i := 0; i < len(bv.Bonuses)-1; i++ {
		// Card rectangle
		x := bonusBgOffsetX + float32(i+1)*bonusCardIntervalOffset + float32(i)*bonusCardWidth
		y := float32(bonusBgOffsetY + bonusCardOffsetY)
		graphics.DrawRect(
			screen,
			x, y,
			bonusCardWidth,
			bonusCardHeight,
			bonusBgColor[0], bonusBgColor[1], bonusBgColor[2], bonusBgColor[3],
		)
		// Card title text
		rect := text.BoundString(assets.CardTitleFontFace, bv.Bonuses[i].Name())
		geom := ebiten.GeoM{}
		geom.Translate(
			float64(x)+float64(bonusCardWidth)/2-float64(rect.Max.X)/2,
			float64(y)+float64(bonusCardHeight)/6-float64(rect.Max.Y)/2,
		)
		text.DrawWithOptions(screen, bv.Bonuses[i].Name(), assets.CardTitleFontFace, &ebiten.DrawImageOptions{
			GeoM: geom,
		})
		// Card description rectangle
		y += bonusCardHeight / 6
		graphics.DrawRect(
			screen,
			x+bonusCardIntervalOffset, y,
			bonusCardWidth-bonusCardIntervalOffset*2,
			bonusCardHeight-bonusCardHeight/6-bonusCardIntervalOffset,
			1, 1, 1, 0.5,
		)

		// Highlight selection
		if i == bv.SelectedIndex {
			graphics.DrawRectBorder(
				screen,
				bonusBgOffsetX+float32(i+1)*bonusCardIntervalOffset+float32(i)*bonusCardWidth,
				bonusBgOffsetY+bonusCardOffsetY,
				bonusCardWidth,
				bonusCardHeight,
				2, 1, 1, 1, 0.9,
			)
		}
	}
	// Skip button
	x := float32(bonusBgOffsetX + bonusCardIntervalOffset)
	y := float32(bonusBgOffsetY + bonusCardOffsetY + bonusCardHeight + bonusCardIntervalOffset)
	graphics.DrawRect(
		screen,
		x, y,
		bonusCardWidth*2+bonusCardIntervalOffset,
		bonusCardHeight/4-bonusCardIntervalOffset,
		0.2, 0.2, 0.2, 0.7,
	)
	// Card title text
	str := "Skip"
	rect := text.BoundString(assets.CardTitleFontFace, str)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(x)+float64(bonusCardWidth)+float64(bonusCardIntervalOffset)/2-float64(rect.Max.X)/2,
		float64(y)+float64(bonusCardHeight)/8-float64(rect.Max.Y)/2,
	)
	text.DrawWithOptions(screen, str, assets.CardTitleFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	// Highlight selection
	if bv.SelectedIndex == len(bv.Bonuses)-1 {
		graphics.DrawRectBorder(
			screen,
			x, y,
			bonusCardWidth*2+bonusCardIntervalOffset,
			bonusCardHeight/4-bonusCardIntervalOffset,
			2, 1, 1, 1, 0.9,
		)
	}
}
