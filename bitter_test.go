package bitter

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAll(t *testing.T) {
	want := []string{"aperol", "campari", "fernet", "cynar"}
	have := make([]string, len(want))

	for i, v := range FromSlice(want).All() {
		have[i] = v
	}

	if !cmp.Equal(want, have) {
		t.Errorf("FromSlice(%q).All() --> %q, want %q", want, have, want)
	}

}

func TestForEach(t *testing.T) {
	in := []string{"aperol", "campari", "fernet", "cynar"}
	want := []string{"APEROL", "CAMPARI", "FERNET", "CYNAR"}

	have := make([]string, len(in))
	doUpper := func(i int, v string) (int, string) {
		return i, strings.ToUpper(v)
	}

	for i, v := range ForEach(FromSlice(want), doUpper).All() {
		have[i] = v
	}

	if !cmp.Equal(want, have) {
		t.Errorf("ForEach(FromSlice(%q), doUpper).All() --> %q, want %q", in, have, want)
	}

}

func TestForEachV(t *testing.T) {
	in := []string{"aperol", "campari", "fernet", "cynar"}
	want := []string{"APEROL", "CAMPARI", "FERNET", "CYNAR"}

	have := make([]string, len(in))
	for i, v := range ForEachV(FromSlice(want), strings.ToUpper).All() {
		have[i] = v
	}

	if !cmp.Equal(want, have) {
		t.Errorf("ForEachV(FromSlice(%q), strings.ToUpper).All() --> %q, want %q", in, have, want)
	}

}
