package ranges

import "golang.org/x/exp/constraints"

type Bound interface {
	boundMarker()
}

type BoundIncluded[T constraints.Ordered] struct {
	Value T
}

func (BoundIncluded[T]) boundMarker() {}

type BoundExcluded[T constraints.Ordered] struct {
	Value T
}

func (BoundExcluded[T]) boundMarker() {}

type BoundUnbounded struct{}

func (BoundUnbounded) boundMarker() {}
