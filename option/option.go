package option

import (
	"fmt"

	xfmt "github.com/FedericoSchonborn/go-sandbox/fmt"
	"github.com/FedericoSchonborn/go-sandbox/result"
	"github.com/FedericoSchonborn/go-sandbox/zero"
)

type Option[T any] struct {
	some  bool
	value T
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		some:  true,
		value: value,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		some: false,
	}
}

func (o Option[T]) Match(someFunc func(T), noneFunc func()) {
	if o.some {
		someFunc(o.value)
	}

	noneFunc()
}

func (o Option[T]) IsSome() bool {
	return o.some
}

func (o Option[T]) IsNone() bool {
	return !o.some
}

func (o Option[T]) Unwrap() T {
	if o.some {
		return o.value
	}

	panic("Called Option.Unwrap on None value")
}

// UnwrapUnchecked unwraps the Option without checking if it is Some, this may
// result on unexpected behavior or a panic.
func (o Option[T]) UnwrapUnchecked() T {
	return o.value
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.some {
		return o.value
	}

	return def
}

func (o Option[T]) UnwrapOrElse(fn func() T) T {
	if o.some {
		return o.value
	}

	return fn()
}

// UnwrapOrZero is unwrap_or_default but more gopher-ish.
func (o Option[T]) UnwrapOrZero() (value T, ok bool) {
	if o.some {
		return o.value, true
	}

	return zero.Zero[T](), false
}

func Map[T, U any](o Option[T], fn func(T) U) Option[U] {
	if o.some {
		return Some(fn(o.value))
	}

	return None[U]()
}

func OkOr[T any, E error](o Option[T], err E) result.Result[T, E] {
	if o.some {
		return result.Ok[T, E](o.value)
	}

	return result.Err[T](err)
}

func OkOrElse[T any, E error](o Option[T], err func() E) result.Result[T, E] {
	if o.some {
		return result.Ok[T, E](o.value)
	}

	return result.Err[T](err())
}

func (o Option[T]) Filter(pred func(T) bool) Option[T] {
	if o.some && pred(o.value) {
		return Some(o.value)
	}

	return None[T]()
}

func (o Option[T]) Or(ob Option[T]) Option[T] {
	if o.some {
		return Some(o.value)
	}

	return ob
}

func (o Option[T]) OrElse(fn func() Option[T]) Option[T] {
	if o.some {
		return Some(o.value)
	}

	return fn()
}

type Zipped[L, R any] struct {
	Left  L
	Right R
}

func Zip[L, R any](l Option[L], r Option[R]) Option[Zipped[L, R]] {
	if l.some && r.some {
		return Some(Zipped[L, R]{
			Left:  l.value,
			Right: r.value,
		})
	}

	return None[Zipped[L, R]]()
}

func (o Option[T]) Format(f fmt.State, verb rune) {
	if o.some {
		fmt.Fprintf(f, "Some("+xfmt.Format(f, verb)+")", o.value)
		return
	}

	fmt.Fprint(f, "None")
}
