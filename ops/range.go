package ops

import "github.com/fdschonborn/go-sandbox/ints"

type RangeBounds[Index ints.Int, Start Bound, End Bound] interface {
	StartBound() Start
	EndBound() End
	Contains(index Index) bool
}

var _ RangeBounds[int, BoundIncluded, BoundExcluded] = Range[int]{}

type Range[Index ints.Int] struct {
	Start, End Index
}

func NewRange[Index ints.Int](start, end Index) Range[Index] {
	return Range[Index]{Start: start, End: end}
}

func (r Range[Index]) StartBound() BoundIncluded {
	return BoundIncluded(r.Start)
}

func (r Range[Index]) EndBound() BoundExcluded {
	return BoundExcluded(r.End)
}

func (r Range[Index]) Contains(index Index) bool {
	return (r.Start <= index) && (index < r.End)
}

var _ RangeBounds[int, BoundIncluded, BoundIncluded] = InclusiveRange[int]{}

type InclusiveRange[Index ints.Int] struct {
	Start, End Index
}

func NewInclusiveRange[Index ints.Int](start, end Index) InclusiveRange[Index] {
	return InclusiveRange[Index]{Start: start, End: end}
}

func (ir InclusiveRange[Index]) StartBound() BoundIncluded {
	return BoundIncluded(ir.Start)
}

func (ir InclusiveRange[Index]) EndBound() BoundIncluded {
	return BoundIncluded(ir.End)
}

func (ir InclusiveRange[Index]) Contains(index Index) bool {
	return (ir.Start <= index) && (index <= ir.End)
}
