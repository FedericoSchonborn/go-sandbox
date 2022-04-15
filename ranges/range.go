package ranges

import "golang.org/x/exp/constraints"

type Interface[T any, S Bound, E Bound] interface {
	StartBound() S
	EndBound() E
	Contains(T) bool
}

var _ Interface[int, BoundIncluded[int], BoundExcluded[int]] = Range[int]{}

type Range[T constraints.Ordered] struct {
	Start, End T
}

func New[T constraints.Ordered](start, end T) Range[T] {
	return Range[T]{Start: start, End: end}
}

func (r Range[T]) StartBound() BoundIncluded[T] {
	return BoundIncluded[T]{r.Start}
}

func (r Range[T]) EndBound() BoundExcluded[T] {
	return BoundExcluded[T]{r.End}
}

func (r Range[T]) Contains(index T) bool {
	return (r.Start <= index) && (index < r.End)
}

var _ Interface[int, BoundIncluded[int], BoundIncluded[int]] = Inclusive[int]{}

type Inclusive[T constraints.Ordered] struct {
	Start, End T
}

func NewInclusive[T constraints.Ordered](start, end T) Inclusive[T] {
	return Inclusive[T]{Start: start, End: end}
}

func (ir Inclusive[T]) StartBound() BoundIncluded[T] {
	return BoundIncluded[T]{ir.Start}
}

func (ir Inclusive[T]) EndBound() BoundIncluded[T] {
	return BoundIncluded[T]{ir.End}
}

func (ir Inclusive[T]) Contains(index T) bool {
	return (ir.Start <= index) && (index <= ir.End)
}
