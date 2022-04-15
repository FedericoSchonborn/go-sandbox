package option

import (
	"fmt"
	"io"

	xfmt "github.com/FedericoSchonborn/go-sandbox/fmt"
	"github.com/FedericoSchonborn/go-sandbox/result"
	"github.com/FedericoSchonborn/go-sandbox/zero"
)

type Option[T any] struct {
	some  bool
	inner T
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		some:  true,
		inner: value,
	}
}

// None and Option{} are equivalent.
func None[T any]() Option[T] {
	return Option[T]{
		some: false,
	}
}

func (o Option[T]) IsSome() bool {
	return o.some
}

func (o Option[T]) IsNone() bool {
	return !o.some
}

func (o Option[T]) Unwrap() T {
	if o.some {
		return o.inner
	}

	panic("Called Option.Unwrap on None value")
}

// UnwrapUnchecked unwraps the Option without checking if it is Some, this may
// result on unexpected behavior or a panic.
func (o Option[T]) UnwrapUnchecked() T {
	return o.inner
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.some {
		return o.inner
	}

	return def
}

type UnwrapOrElseFunc[T any] func() T

func (o Option[T]) UnwrapOrElse(fn UnwrapOrElseFunc[T]) T {
	if o.some {
		return o.inner
	}

	return fn()
}

// UnwrapOrZero is unwrap_or_default but more gopher-ish.
func (o Option[T]) UnwrapOrZero() (value T, ok bool) {
	if o.some {
		return o.inner, true
	}

	return zero.Zero[T](), false
}

type MapFunc[T, U any] func(T) U

func Map[T, U any](o Option[T], fn MapFunc[T, U]) Option[U] {
	if o.some {
		return Some(fn(o.inner))
	}

	return None[U]()
}

func OkOr[T any, E error](o Option[T], err E) result.Result[T] {
	if o.some {
		return result.Ok(o.inner)
	}

	return result.Err[T](err)
}

type OkOrElseFunc[E error] func() E

func OkOrElse[T any, E error](o Option[T], err OkOrElseFunc[E]) result.Result[T] {
	if o.some {
		return result.Ok(o.inner)
	}

	return result.Err[T](err())
}

type FilterFunc[T any] func(T) bool

func (o Option[T]) Filter(pred FilterFunc[T]) Option[T] {
	if o.some && pred(o.inner) {
		return Some(o.inner)
	}

	return None[T]()
}

func (o Option[T]) Or(ob Option[T]) Option[T] {
	if o.some {
		return Some(o.inner)
	}

	return ob
}

type OrElseFunc[T any] func() Option[T]

func (o Option[T]) OrElse(fn OrElseFunc[T]) Option[T] {
	if o.some {
		return Some(o.inner)
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
			Left:  l.inner,
			Right: r.inner,
		})
	}

	return None[Zipped[L, R]]()
}

func (o Option[T]) Format(f fmt.State, verb rune) {
	if !o.some {
		io.WriteString(f, "None")
		return
	}

	fmt.Fprintf(f, "Some("+xfmt.Format(f, verb)+")", o.inner)
}
