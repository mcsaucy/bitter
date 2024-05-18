// Package bitter lets you build pipelines using rangefunc sequences.
package bitter

import (
	"context"
	"iter"
)

func FromSlice[V any](s []V) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

// ForEach executes "do" for each entry of "in".
func ForEach[K, V, Kp, Vp any](in iter.Seq2[K, V], do func(K, V) (Kp, Vp)) iter.Seq2[Kp, Vp] {
	return func(yield func(Kp, Vp) bool) {
		for k, v := range in {
			if !yield(do(k, v)) {
				return
			}
		}
	}
}

// ForEachV is like ForEach, except it only processes the right value (V) and preserves the left value (K).
func ForEachV[K, V, Vp any](in iter.Seq2[K, V], do func(V) Vp) iter.Seq2[K, Vp] {
	return ForEach(in, func(k K, v V) (K, Vp) {
		return k, do(v)
	})
}

// ForEachContext executes "do" for each entry of "in", passing in a context as the first argument.
func ForEachContext[K, V, Kp, Vp any](
	ctx context.Context,
	in iter.Seq2[K, V],
	do func(context.Context, K, V) (Kp, Vp),
) iter.Seq2[Kp, Vp] {
	return func(yield func(Kp, Vp) bool) {
		for k, v := range in {
			if !yield(do(ctx, k, v)) {
				return
			}
		}
	}
}

// ForEachVContext is like ForEach, except it only processes the right value (V) and preserves the left value (K), passing in a context as the first argument.
func ForEachVContext[K, V, Vp any](
	ctx context.Context,
	in iter.Seq2[K, V],
	do func(context.Context, V) Vp,
) iter.Seq2[K, Vp] {
	return ForEachContext(ctx, in, func(ctx context.Context, k K, v V) (K, Vp) {
		return k, do(ctx, v)
	})
}
