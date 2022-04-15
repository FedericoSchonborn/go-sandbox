package randx

import "math/rand"

type Rand struct {
	*rand.Rand
}

func New(src rand.Source) *Rand {
	return &Rand{
		Rand: rand.New(src),
	}
}

func (rx *Rand) Bool() bool {
	return rx.Intn(2) == 0
}

func Bool() bool {
	return rand.Intn(2) == 0
}
