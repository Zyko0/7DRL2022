package main

import (
	"errors"
	"fmt"

	"github.com/Zyko0/7DRL2022/core"
	"github.com/Zyko0/7DRL2022/core/bonus"
	"github.com/Zyko0/7DRL2022/graphics"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/Zyko0/7DRL2022/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	core     *core.Core
	renderer *graphics.Renderer

	bonusView *ui.BonusView
}

func New() *Game {
	c := core.NewCore()
	c.Initialize()

	return &Game{
		core:     c,
		renderer: graphics.NewRenderer(),

		bonusView: ui.NewBonusView(),
	}
}

func (g *Game) Update() error {
	// TODO: remove
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("quit")
	}

	// If needs an augment selection but the view is not active yet, roll, activate and abort
	if g.core.ChestPickedUp {
		if !g.bonusView.Active() {
			primary, secondary := g.core.BonusList.Roll()
			g.bonusView.SetBonuses([]bonus.Bonus{
				primary, secondary, bonus.BonusNone,
			})
			return nil
		}
		// Update and check for user input, if an augment is picked, the view isn't active anymore
		g.bonusView.Update()
		if g.bonusView.Active() {
			// Abort because there's still an augment to pick
			return nil
		}
		// If the view is not active anymore, check for selection
		b := g.bonusView.Bonuses[g.bonusView.SelectedIndex]
		// Otherwise pick augment
		g.core.BonusList.Consume(b)
		g.core.Stats.ApplyBonus(b)
		g.core.ChestPickedUp = false
		g.bonusView.Reset()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.core = core.NewCore()
		g.core.Initialize()
	}

	g.core.Update()
	g.renderer.Update(g.core.Player.Y)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.RenderBackground(screen)
	g.renderer.RenderPlatforms(screen, g.core.Platforms)
	g.renderer.RenderEntities(screen, g.core.Entities)
	g.renderer.RenderPlayer(screen, g.core.Player)
	g.renderer.RenderWave(screen, g.core.Wave)
	g.renderer.RenderHUD(screen, g.core.Player.HP, uint64(g.core.Player.Height))
	// Main menu view with background demo
	if g.bonusView.Active() {
		g.bonusView.Draw(screen)
		return
	}
	// Debug
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f - FPS %.02f PlayerPos [%.2f, %.2f] - Platforms %d",
			ebiten.CurrentTPS(),
			ebiten.CurrentFPS(),
			g.core.Player.X, g.core.Player.Y,
			g.core.Platforms.Count(),
		),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return logic.ScreenWidth, logic.ScreenHeight
}

func main() {
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum) // TODO: vsync on
	ebiten.SetMaxTPS(logic.TPS)
	ebiten.SetFullscreen(true)
	ebiten.SetCursorMode(ebiten.CursorModeCaptured)

	if err := ebiten.RunGame(New()); err != nil {
		fmt.Println("rungame:", err)
	}
}
