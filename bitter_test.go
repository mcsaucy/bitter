package bitter

import (
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
