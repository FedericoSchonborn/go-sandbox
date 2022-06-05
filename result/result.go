//go:build go1.18

package result

import (
	"fmt"

	xfmt "github.com/FedericoSchonborn/go-sandbox/fmt"
)

type Result[T, E any] struct {
	ok    bool
	inner any // T | E
}

func Ok[T, E any](value T) Result[T, E] {
	return Result[T, E]{
		ok:    true,
		inner: value,
	}
}

func Err[T, E any](err E) Result[T, E] {
	return Result[T, E]{
		ok:    false,
		inner: err,
	}
}

func (r Result[T, E]) Match(okFunc func(T), errFunc func(E)) {
	if r.ok {
		okFunc(r.inner.(T))
	}

	errFunc(r.inner.(E))
}

func (r Result[T, E]) IsOk() bool {
	return r.ok
}

func (r Result[T, E]) IsErr() bool {
	return !r.ok
}

func Map[T, E, U any](r Result[T, E], op func(T) U) Result[U, E] {
	if !r.ok {
		return Err[U](r.inner.(E))
	}

	return Ok[U, E](op(r.inner.(T)))
}

func (r Result[T, E]) Format(f fmt.State, verb rune) {
	var tag string
	if r.ok {
		tag = "Ok("
	} else {
		tag = "Err("
	}

	fmt.Fprintf(f, tag+xfmt.Format(f, verb)+")", r.inner)
}
