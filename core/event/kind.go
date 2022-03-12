package event

type Kind byte

const (
	KindNone Kind = iota
	KindAoeSpawn
	KindEnemySpawn
	KindChestSpawn
	KindReduceMaxPlatformWidth
	KindEnemyUpgrade
	KindEnemyFaster
)
