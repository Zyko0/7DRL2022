package ui

import (
	"fmt"

	"github.com/Zyko0/7DRL2022/assets"
	"github.com/Zyko0/7DRL2022/graphics"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	gameoverCardWidth   = 0.8 * logic.ScreenHeight
	gameoverCardHeight  = 0.2 * logic.ScreenHeight
	gameoverCardOffsetX = (logic.ScreenWidth - gameoverCardWidth) / 2
	gameoverCardOffsetY = (logic.ScreenHeight - gameoverCardHeight) / 2
)

type GameoverView struct {
	active bool

	score uint64

	card *ebiten.Image
}

func NewGameoverView() *GameoverView {
	card := ebiten.NewImage(gameoverCardWidth, gameoverCardHeight)
	graphics.DrawRect(card, 0, 0, gameoverCardWidth, gameoverCardHeight, 0.1, 0.1, 0.1, 0.9)
	graphics.DrawRectBorder(card, 0, 0, gameoverCardWidth, gameoverCardHeight, 1, 1, 1, 1, 0.9)
	// Title
	str := "Game over"
	rect := text.BoundString(assets.CardTitleFontFace, str)
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(gameoverCardWidth/2-rect.Max.X/2),
		float64(48),
	)
	colorM := ebiten.ColorM{}
	colorM.Scale(0.8, 0, 0, 1)
	text.DrawWithOptions(card, str, assets.CardTitleFontFace, &ebiten.DrawImageOptions{
		GeoM:   geom,
		ColorM: colorM,
	})
	// Resume text
	str = "Press 'R' to start a new game"
	rect = text.BoundString(assets.CardTitleFontFace, str)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(gameoverCardWidth/2-rect.Max.X/2),
		float64(96),
	)
	text.DrawWithOptions(card, str, assets.CardTitleFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	return &GameoverView{
		score: 0,

		card: card,
	}
}

func (gv *GameoverView) Active() bool {
	return gv.active
}

func (gv *GameoverView) Update(isGameOver bool, score uint64) {
	gv.active = isGameOver
	gv.score = score
	if gv.active {
		/*assets.StopInGameMusic()
		assets.PlayGameoverMusic()*/
	} else {
		// assets.StopGameoverMusic()
	}
}

func (gv *GameoverView) Draw(screen *ebiten.Image) {
	geom := ebiten.GeoM{}
	geom.Translate(
		float64(gameoverCardOffsetX),
		float64(gameoverCardOffsetY),
	)
	screen.DrawImage(gv.card, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
	str := fmt.Sprintf("Best height: %d", gv.score)
	rect := text.BoundString(assets.CardTitleFontFace, str)
	geom = ebiten.GeoM{}
	geom.Translate(
		float64(gameoverCardOffsetX+gameoverCardWidth/2-rect.Max.X/2),
		float64(gameoverCardOffsetY+gameoverCardHeight/2-rect.Max.Y/2+32),
	)
	text.DrawWithOptions(screen, str, assets.CardTitleFontFace, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}
