//go:build go1.18

package result

import (
	"fmt"

	xfmt "github.com/FedericoSchonborn/go-sandbox/fmt"
)

type Result[T any] struct {
	ok    bool
	inner any // T | error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		ok:    true,
		inner: value,
	}
}

// Result{} and Err(nil) are equivalent.
func Err[T any](err error) Result[T] {
	return Result[T]{
		ok:    false,
		inner: err,
	}
}

func (r Result[T]) IsOk() bool {
	return r.ok
}

func (r Result[T]) IsErr() bool {
	return !r.ok
}

func Map[T any, U any](r Result[T], op func(T) U) Result[U] {
	if !r.ok {
		return Err[U](r.inner.(error))
	}

	return Ok(op(r.inner.(T)))
}

func (r Result[T]) Format(f fmt.State, verb rune) {
	var kind string
	if r.ok {
		kind = "Ok("
	} else {
		kind = "Err("
	}

	fmt.Fprintf(f, kind+xfmt.Format(f, verb)+")", r.inner)
}
