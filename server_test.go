package gopher

import (
	"strings"
	"testing"
)

func TestServerErrorListing(t *testing.T) {
	s := &Server{}
	l := s.ErrorListing("test error")

	if got, want := l.String(), "\x03test error\t\t\t\r\n"; got != want {
		t.Fatalf(`s.ErrorListing("test error").String() = %q, want %q`, got, want)
	}
}

func TestServerFilename(t *testing.T) {
	for i, tt := range []struct {
		root string
		path string
		want string
	}{
		{"/foo/bar", "baz/qux", "/foo/bar/baz/qux"},
		{"/baz/qux", "sit/./amet", "/baz/qux/sit/amet"},
	} {
		s := &Server{Root: tt.root}

		if got := s.filename(tt.path); got != tt.want {
			t.Fatalf(`[%d] s.filename(%q) = %q, want %q`, i, tt.path, got, tt.want)
		}
	}
}

func TestGetPath(t *testing.T) {
	for i, tt := range []struct {
		in   string
		want string
	}{
		{"foo/bar\nbaz", "foo/bar"},
		{"baz\nqux", "baz"},
	} {
		s := &Server{}

		path, err := s.getPath(strings.NewReader(tt.in))
		if err != nil {
			t.Fatalf(`[%d] err = %v`, i, err)
		}

		if path != tt.want {
			t.Fatalf(`path = %q, want %q`, path, tt.want)
		}
	}
}
