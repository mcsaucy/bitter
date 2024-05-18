// Package bitter lets you build pipelines using rangefunc sequences.
package bitter

import (
	"iter"
)

// Iterator exposes helpers for turning a rangefunc sequence into a pipeline.
type Iterator[K, V any] struct {
	s iter.Seq2[K, V]
}

func FromSlice[V any](s []V) Iterator[int, V] {
	return Iterator[int, V]{
			s: func(yield func(int, V) bool) {
			for i, v := range s {
				if !yield(i, v) {
					return
				}
			}
		},
	}
}

// All exposes an iter.Seq for use with range or iter.Pull.
func (i Iterator[K, V]) All() iter.Seq2[K, V] {
	return i.s
}
