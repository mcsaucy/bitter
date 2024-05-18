package bitter

import (
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAll(t *testing.T) {
	want := []string{"aperol", "campari", "fernet", "cynar"}
	have := make([]string, len(want))

	for i, v := range FromSlice2(want) {
		have[i] = v
	}

	if !cmp.Equal(want, have) {
		t.Errorf("FromSlice2(%q) --> %q, want %q", want, have, want)
	}

}

func TestForEach(t *testing.T) {
	in := []string{"aperol", "campari", "fernet", "cynar"}
	want := []string{"APEROL", "CAMPARI", "FERNET", "CYNAR"}

	doUpper := func(i int, v string) (int, string) {
		return i, strings.ToUpper(v)
	}

	have := ToSlice(JustV(ForEach2(FromSlice2(want), doUpper)))

	if !cmp.Equal(want, have) {
		t.Errorf("ForEach2(FromSlice2(%q), doUpper) --> %q, want %q", in, have, want)
	}

}

func TestForEachJustV(t *testing.T) {
	in := []string{"aperol", "campari", "fernet", "cynar"}
	want := []string{"APEROL", "CAMPARI", "FERNET", "CYNAR"}

	have := ToSlice(ForEach(JustV(FromSlice2(want)), strings.ToUpper))

	if !cmp.Equal(want, have) {
		t.Errorf("ForEach(FromSlice(%q), strings.ToUpper) --> %q, want %q", in, have, want)
	}

}

func TestForEachVContext(t *testing.T) {
	in := []string{"aperol", "campari", "fernet", "cynar"}
	want := []string{"APEROL", "CAMPARI", "FERNET", "CYNAR"}

	ctx := context.WithValue(context.Background(), "lol", "omg")
	do := func(ctx context.Context, in string) string {
		v := ctx.Value("lol")
		if vs, ok := v.(string); !ok || vs != "omg" {
			t.Errorf("context.Value(lol) = %v; want omg", v)
		}
		return strings.ToUpper(in)
	}

	have := ToSlice(ForEachContext(ctx, JustV(FromSlice2(want)), do))

	if !cmp.Equal(want, have) {
		t.Errorf("ForEachVContext(ctx, FromSlice(%q), strings.ToUpper) --> %q, want %q", in, have, want)
	}

}
