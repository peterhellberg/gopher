package gopher

import "testing"

func TestSuffixes(t *testing.T) {
	if got, want := len(Suffixes), 17; got != want {
		t.Fatalf(`len(Suffixes) = %d, want %d`, got, want)
	}
}
