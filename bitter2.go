package bitter

import (
	"context"
	"iter"
)

func FromSlice2[V any](s []V) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		for i := range s {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

// ForEach2 executes "do" for each entry of "in".
func ForEach2[K, V, Kp, Vp any](in iter.Seq2[K, V], do func(K, V) (Kp, Vp)) iter.Seq2[Kp, Vp] {
	return func(yield func(Kp, Vp) bool) {
		for k, v := range in {
			if !yield(do(k, v)) {
				return
			}
		}
	}
}

// JustK turns a Seq2[K,V] and emits a Seq[K]
func JustK[K, V any](in iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(k K) bool) {
		for k, _ := range in {
			if !yield(k) {
				return
			}
		}
	}
}

// JustK turns a Seq2[K,V] and emits a Seq[V]
func JustV[K, V any](in iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(v V) bool) {
		for _, v := range in {
			if !yield(v) {
				return
			}
		}
	}
}

// ForEachContext executes "do" for each entry of "in", passing in a context as the first argument.
func ForEachContext2[K, V, Kp, Vp any](
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
