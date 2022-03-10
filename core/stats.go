package core

type Stats struct {
}

func (s *Stats) Clone() *Stats {
	cloned := &Stats{}
	*cloned = *s

	return cloned
}
