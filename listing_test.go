package gopher

import "testing"

func TestSuffixes(t *testing.T) {
	if got, want := len(Suffixes), 17; got != want {
		t.Fatalf(`len(Suffixes) = %d, want %d`, got, want)
	}
}

func TestListingString(t *testing.T) {
	for _, tt := range []struct {
		l    Listing
		want string
	}{
		{Listing{}, ""},
		{Listing{[]Entry{
			{'0', "Foo", "foo", "test", "70"},
			{'1', "Bar", "bar", "test", "70"},
		}}, "1Bar\tbar\ttest\t70\r\n0Foo\tfoo\ttest\t70\r\n"},
		{Listing{[]Entry{
			{'I', "Baz", "baz", "test2", "7070"},
			{'1', "Qux", "qux", "test2", "7070"},
		}}, "1Qux\tqux\ttest2\t7070\r\nIBaz\tbaz\ttest2\t7070\r\n"},
	} {
		if got := tt.l.String(); got != tt.want {
			t.Fatalf(`l.String() = %q, want %q`, got, tt.want)
		}
	}
}

func TestListingVisitDir(t *testing.T) {
	l := Listing{[]Entry{
		{'0', "Foo", "foo", "test", "70"},
		{'1', "Bar", "bar", "test", "70"},
	}}

	l.VisitDir("Dir", "bar/baz", "bar/", "example", "70")

	if got, want := len(l.entries), 3; got != want {
		t.Fatalf(`len(l.entries) = %d, want %d`, got, want)
	}

	e := l.entries[2]

	if got, want := e.String(), "1Dir\tbaz\texample\t70\r\n"; got != want {
		t.Fatalf(`e.String() = %q, want %q`, got, want)
	}
}

func TestListingVisitFile(t *testing.T) {
	l := Listing{[]Entry{
		{'0', "Foo", "foo", "test", "70"},
		{'1', "Bar", "bar", "test", "70"},
	}}

	l.VisitFile("File", "bar/baz.png", "bar/", "example", "70")

	if got, want := len(l.entries), 3; got != want {
		t.Fatalf(`len(l.entries) = %d, want %d`, got, want)
	}

	e := l.entries[2]

	if got, want := e.String(), "IFile\tbaz.png\texample\t70\r\n"; got != want {
		t.Fatalf(`e.String() = %q, want %q`, got, want)
	}
}
