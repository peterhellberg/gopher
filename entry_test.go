package gopher

import "testing"

func TestEntryString(t *testing.T) {
	for _, tt := range []struct {
		e    Entry
		want string
	}{
		{Entry{'0', "a", "b", "c", "d"}, "0a\tb\tc\td\r\n"},
		{Entry{'9', "e", "f", "g", "h"}, "9e\tf\tg\th\r\n"},
	} {
		if got := tt.e.String(); got != tt.want {
			t.Fatalf(`tt.e.String() = %q, want %q`, got, tt.want)
		}
	}
}
