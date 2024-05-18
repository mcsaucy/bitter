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

// ForEach executes "do" for each entry of "in".
func ForEach[K, V, Kp, Vp any](in Iterator[K, V], do func(K, V) (Kp, Vp)) Iterator[Kp, Vp] {
	return Iterator[Kp, Vp]{
		s: func(yield func(Kp, Vp) bool) {
			for k, v := range in.s {
				kp, vp := do(k, v)
				if !yield(kp, vp) {
					return
				}
			}
		},
	}
}

// ForEachV is like ForEach, except it only processes the right value (V) and preserves the left value (K).
func ForEachV[K, V, Vp any](in Iterator[K, V], do func(V) Vp) Iterator[K, Vp] {
	return ForEach(in, func(k K, v V) (K, Vp) {
		return k, do(v)
	})
}
