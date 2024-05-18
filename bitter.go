// Package bitter lets you build cool rangefuncs using iter.Seq and iter.Seq2.
package bitter

import (
	"context"
	"iter"
)

func FromSlice[V any](s []any) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range s {
			if !yield(i) {
				return
			}
		}
	}
}

func ToSlice[V any](i iter.Seq[V]) []V {
	var vs []V
	for v := range i {
		vs = append(vs, v)
	}
	return vs
}

// ForEach executes "do" for each entry of "in".
func ForEach[V, Vp any](in iter.Seq[V], do func(V) Vp) iter.Seq[Vp] {
	return func(yield func(v Vp) bool) {
		for v := range in {
			if !yield(do(v)) {
				return
			}
		}
	}
}

// Enhance executes "do" for each entry of "in", expecting "do" to return two values.
func Enhance[V, Kp, Vp any](in iter.Seq[V], do func(V) (Kp, Vp)) iter.Seq2[Kp, Vp] {
	return func(yield func(k Kp, v Vp) bool) {
		for v := range in {
			if !yield(do(v)) {
				return
			}
		}
	}
}

// ForEachContext is like ForEach, but also passes in a context as the first argument.
func ForEachContext[V, Vp any](
	ctx context.Context,
	in iter.Seq[V],
	do func(context.Context, V) Vp,
) iter.Seq[Vp] {
	return ForEach(in, func(v V) Vp {
		return do(ctx, v)
	})
}
