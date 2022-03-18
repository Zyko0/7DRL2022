package core

import (
	"math/rand"
	"time"

	"github.com/Zyko0/7DRL2022/assets"
	"github.com/Zyko0/7DRL2022/core/bonus"
	"github.com/Zyko0/7DRL2022/core/entity"
	"github.com/Zyko0/7DRL2022/core/event"
	"github.com/Zyko0/7DRL2022/core/platform"
	"github.com/Zyko0/7DRL2022/core/utils"
	"github.com/Zyko0/7DRL2022/logic"
)

const (
	ColumnsCount = logic.ScreenWidth / logic.UnitSize

	WaveIncreaseFrequencyHeightInterval = 100.
)

type Core struct {
	ticks        uint64
	rng          *rand.Rand
	nextHeight   float64
	eventManager *event.Manager

	bestHeight        float64
	lastChestPlatform *platform.Platform

	nextCheckpointTick uint64

	ChestPickedUp bool

	BonusList *bonus.List
	Stats     *Stats
	Wave      *Wave
	Platforms *platform.List
	Entities  []entity.Entity
	Player    *Player
}

func NewCore() *Core {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	return &Core{
		ticks:        0,
		rng:          rng,
		nextHeight:   WaveIncreaseFrequencyHeightInterval,
		eventManager: event.NewManager(),

		nextCheckpointTick: 0,

		ChestPickedUp: false,

		BonusList: bonus.NewList(rng),
		Stats:     NewStats(),
		Wave:      NewWave(rng),
		Platforms: platform.NewList(),
		Entities:  []entity.Entity{},
		Player:    NewPlayer(),
	}
}

func (c *Core) Initialize() {
	c.Platforms.Initialize(c.rng)
	// Set the player on the first platform
	p := c.Platforms.Slice()[0]
	c.Player.X = p.X - logic.UnitSize/2
	c.Player.Y = p.Y + logic.UnitSize/2 + c.Player.Height/2
	c.Player.GroundedPlatform = p
}

func (c *Core) TriggerChestPlatformWave() {
	c.lastChestPlatform.Enlarge()
	c.Wave.reach(c.Player.Y - logic.ScreenHeight/2)
	if c.Stats.CheckpointMod > 0 {
		c.nextCheckpointTick = c.ticks + c.Stats.CheckpointMod
	}
}

func (c *Core) handlePlatformGeneration() {
	p := c.Platforms.GetHighestPlatform()

	missingRange := p.Y - c.Player.Y - logic.ScreenHeight/2
	jumpv := utils.JumpVector(c.Stats.JumpForce, 1)
	jumpv[0] *= c.Player.MoveSpeed
	jumpv[1] *= c.Player.MoveSpeed
	// Generate missing platforms in advance
	for missingRange < 0 {
		count := c.Stats.PlatformCellCount + c.rng.Intn(platform.BaseCellsCount-c.Stats.PlatformCellCount+1)
		p = platform.Generate(c.rng, p, count, jumpv)
		c.Platforms.AddPlatform(p)
		missingRange = p.Y - c.Player.Y - logic.ScreenHeight/2
	}
}

func (c *Core) handleVelocity() {
	// TODO: Or not grounded but double jump allowed
	if c.Player.GroundedPlatform != nil && c.Player.IntentVector[1] == 1 {
		c.Player.GroundedPlatform = nil
		jv := utils.JumpVector(c.Stats.JumpForce, c.Player.Orientation)
		c.Player.VelocityVector[0] = jv[0] * c.Player.MoveSpeed
		c.Player.VelocityVector[1] = jv[1] * c.Player.MoveSpeed
		assets.PlayJumpSound()
	}
	// If not in air and not jumping
	if c.Player.GroundedPlatform != nil && c.Player.VelocityVector[1] == 0 {
		c.Player.VelocityVector[0] = c.Player.IntentVector[0] * c.Player.MoveSpeed
	} else if c.Stats.AirControl && c.Player.IntentVector[0] != 0 {
		c.Player.VelocityVector[0] = c.Player.IntentVector[0] * c.Player.MoveSpeed
	}
}

func (c *Core) handleCollisions() {
	y0 := c.Player.Y - logic.ScreenHeight/2
	y1 := y0 + logic.ScreenHeight
	nx := c.Player.X + c.Player.VelocityVector[0]
	ny := c.Player.Y + c.Player.VelocityVector[1]
	if c.Player.VelocityVector[1] < 0 {
		platforms := c.Platforms.AppendPlatformsInRange(nil, y0, y1)
		playerX0 := c.Player.X - c.Player.Width/2
		playerX1 := c.Player.X + c.Player.Width/2
		for _, p := range platforms {
			px0 := p.X - p.Width/2
			px1 := p.X + p.Width/2
			if (playerX0 >= px0 && playerX0 <= px1) || (playerX1 >= px0 && playerX1 <= px1) {
				py1 := p.Y + logic.UnitSize/2
				ny0 := ny - c.Player.Height/2
				if ny0 <= py1 && c.Player.Y-c.Player.Height/2 >= py1 {
					dy := ny - p.GroundY
					ny = p.GroundY
					if c.Player.VelocityVector[0] != 0 {
						dx := dy / c.Player.VelocityVector[1]
						nx = c.Player.X + dx
					}
					c.Player.GroundedPlatform = p
					c.Player.VelocityVector[0] = 0
					c.Player.VelocityVector[1] = 0
					break
				}
			}
		}
	} else if p := c.Player.GroundedPlatform; p != nil && c.Player.VelocityVector[0] != 0 {
		nx0 := nx - c.Player.Width/2
		nx1 := nx + c.Player.Width/2
		// Unground player if out of the platform
		if (nx0 < p.X-p.Width/2 || nx0 > p.X+p.Width/2) && (nx1 < p.X-p.Width/2 || nx1 > p.X+p.Width/2) {
			c.Player.GroundedPlatform = nil
		}
	}
	// Keep player on screen
	if nx < c.Player.Width/2 {
		nx = c.Player.Width / 2
	}
	if nx > logic.ScreenWidth-c.Player.Width/2 {
		nx = logic.ScreenWidth - c.Player.Width/2
	}
	// Set new position
	c.Player.X = nx
	c.Player.Y = ny
	// Resolve fall speed
	if c.Player.GroundedPlatform == nil {
		c.Player.VelocityVector[1] += utils.FallSpeed
		if c.Player.VelocityVector[1] < utils.MaxFallSpeed {
			c.Player.VelocityVector[1] = utils.MaxFallSpeed
		}
	}
}

func (c *Core) Update() {
	if c.Player.Y > c.bestHeight {
		c.bestHeight = c.Player.Y
	}

	// Events
	c.eventManager.Update(c.Player.Y)
	c.handleEvents()
	// Platforms
	c.handlePlatformGeneration()
	if c.Stats.CheckpointMod > 0 && c.ticks > c.nextCheckpointTick {
		if p := c.Player.GroundedPlatform; p != nil {
			c.TriggerChestPlatformWave()
		}
	}
	// Player
	c.Player.Update()
	c.handleVelocity()
	c.handleCollisions()
	// Entities
	c.handleEntities()

	// Wave
	c.Wave.Update(c.Player, c.Stats.WaveHealMod)

	c.ticks++
}

func (c *Core) GetHeight() uint64 {
	return uint64(c.Player.Y) - logic.ScreenHeight/2 - logic.UnitSize*2
}

func (c *Core) GetBestHeight() uint64 {
	return uint64(c.bestHeight) - logic.ScreenHeight/2 - logic.UnitSize*2
}

func (c *Core) IsGameOver() bool {
	return c.Player.HP <= 0
}
