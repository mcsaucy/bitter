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

	for i, v := range FromSlice(want) {
		have[i] = v
	}

	if !cmp.Equal(want, have) {
		t.Errorf("FromSlice(%q) --> %q, want %q", want, have, want)
	}

}

func TestForEach(t *testing.T) {
	in := []string{"aperol", "campari", "fernet", "cynar"}
	want := []string{"APEROL", "CAMPARI", "FERNET", "CYNAR"}

	have := make([]string, len(in))
	doUpper := func(i int, v string) (int, string) {
		return i, strings.ToUpper(v)
	}

	for i, v := range ForEach(FromSlice(want), doUpper) {
		have[i] = v
	}

	if !cmp.Equal(want, have) {
		t.Errorf("ForEach(FromSlice(%q), doUpper) --> %q, want %q", in, have, want)
	}

}

func TestForEachV(t *testing.T) {
	in := []string{"aperol", "campari", "fernet", "cynar"}
	want := []string{"APEROL", "CAMPARI", "FERNET", "CYNAR"}

	have := make([]string, len(in))
	for i, v := range ForEachV(FromSlice(want), strings.ToUpper) {
		have[i] = v
	}

	if !cmp.Equal(want, have) {
		t.Errorf("ForEachV(FromSlice(%q), strings.ToUpper) --> %q, want %q", in, have, want)
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

	have := make([]string, len(in))
	for i, v := range ForEachVContext(ctx, FromSlice(want), do) {
		have[i] = v
	}

	if !cmp.Equal(want, have) {
		t.Errorf("ForEachVContext(ctx, FromSlice(%q), strings.ToUpper) --> %q, want %q", in, have, want)
	}

}
