package either

import (
	"fmt"

	xfmt "github.com/FedericoSchonborn/go-sandbox/fmt"
)

type Either[L, R any] struct {
	left  bool
	inner any // L | R
}

func Left[L, R any](value L) Either[L, R] {
	return Either[L, R]{
		left:  true,
		inner: value,
	}
}

func Right[L, R any](value R) Either[L, R] {
	return Either[L, R]{
		left:  false,
		inner: value,
	}
}

func (e Either[L, R]) Match(leftFunc func(L), rightFunc func(R)) {
	if e.left {
		leftFunc(e.inner.(L))
	}

	rightFunc(e.inner.(R))
}

func (e Either[L, R]) IsLeft() bool {
	return e.left
}

func (e Either[L, R]) IsRight() bool {
	return !e.left
}

func (e Either[L, R]) Format(f fmt.State, verb rune) {
	var kind string
	if e.left {
		kind = "Left("
	} else {
		kind = "Right("
	}

	fmt.Fprintf(f, kind+xfmt.Format(f, verb)+")", e.inner)
}
