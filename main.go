package main

import (
	"errors"
	"fmt"

	"github.com/Zyko0/7DRL2022/core"
	"github.com/Zyko0/7DRL2022/graphics"
	"github.com/Zyko0/7DRL2022/logic"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	core     *core.Core
	renderer *graphics.Renderer
}

func New() *Game {
	c := core.NewCore()
	c.Initialize()

	return &Game{
		core:     c,
		renderer: graphics.NewRenderer(),
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("quit")
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
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetMaxTPS(logic.TPS)
	ebiten.SetFullscreen(true)
	ebiten.SetCursorMode(ebiten.CursorModeCaptured)

	if err := ebiten.RunGame(New()); err != nil {
		fmt.Println("rungame:", err)
	}
}
