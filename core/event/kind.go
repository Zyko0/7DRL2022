package event

type Kind byte

const (
	KindNone Kind = iota
	KindAoeSpawn
	KindEnemySpawn
	KindWaveReset
	KindChestSpawn
	KindSpecialPlatforms
	KindReducePlatformWidth
	KindEnemyUpgrade
	KindEnemyFaster
)
