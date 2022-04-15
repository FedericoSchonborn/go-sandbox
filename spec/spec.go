package spec

import "reflect"

type Spec[T comparable] struct {
	inner T

	Should Should[T]
}

func Value[T comparable](value T) Spec[T] {
	return Spec[T]{
		inner: value,

		Should: Should[T]{
			Equal: func(other T) bool {
				return value == other
			},

			Be: ShouldBe[T]{
				True: func() bool {
					return reflect.ValueOf(value).Bool()
				},
				False: func() bool {
					return !reflect.ValueOf(value).Bool()
				},
			},

			Not: ShouldNot[T]{
				Equal: func(other T) bool {
					return value != other
				},

				Be: ShouldNotBe[T]{
					True: func() bool {
						return !reflect.ValueOf(value).Bool()
					},
					False: func() bool {
						return reflect.ValueOf(value).Bool()
					},
				},
			},
		},
	}
}

type Should[T comparable] struct {
	Equal func(other T) bool
	Be    ShouldBe[T]
	Not   ShouldNot[T]
}

type ShouldBe[T comparable] struct {
	True  func() bool
	False func() bool
}

type ShouldNot[T comparable] struct {
	Equal func(other T) bool
	Be    ShouldNotBe[T]
}

type ShouldNotBe[T comparable] struct {
	True  func() bool
	False func() bool
}
