package gopher

import "testing"

func TestEntryString(t *testing.T) {
	for _, tt := range []struct {
		e    Entry
		want string
	}{
		{Entry{'0', "a", "b", "c", 70}, "0a\tb\tc\t70\r\n"},
		{Entry{'9', "e", "f", "g", 7070}, "9e\tf\tg\t7070\r\n"},
	} {
		if got := tt.e.String(); got != tt.want {
			t.Fatalf(`tt.e.String() = %q, want %q`, got, tt.want)
		}
	}
}
