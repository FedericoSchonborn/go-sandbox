//go:build go1.18

package task

import (
	"sync"
)

type Handle[T any] struct {
	done  chan bool
	value T
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
		done: make(chan bool),
	}

	go func() {
		handle.value = fn()
		handle.done <- true
	}()
	return handle
}

func (h *Handle[T]) Join() T {
	<-h.done
	return h.value
}
