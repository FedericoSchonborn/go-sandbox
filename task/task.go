//go:build go1.18

package task

import (
	"sync"
)

type Handle[T any] struct {
	done  bool
	value chan T
}

func Join[T any](handles ...*Handle[T]) []T {
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		values []T
	)

	for _, handle := range handles {
		wg.Add(1)

		go func(handle *Handle[T]) {
			defer wg.Done()
			value := handle.Join()
			mu.Lock()
			values = append(values, value)
			mu.Unlock()
		}(handle)
	}

	wg.Wait()
	return values
}

func New[T any](fn func() T) *Handle[T] {
	handle := &Handle[T]{
		value: make(chan T, 1),
	}

	go func() {
		handle.value <- fn()
	}()
	return handle
}

func (h *Handle[T]) Join() T {
	if h.done {
		panic("attempt to join an already completed task")
	}

	value := <-h.value
	h.done = true
	return value
}
