package gopher

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
)

var Suffixes = map[string]byte{
	"aiff":     's',
	"au":       's',
	"gif":      'g',
	"go":       '0',
	"gpg":      '0',
	"html":     'h',
	"jpeg":     'I',
	"jpg":      'I',
	"json":     '0',
	"log":      '0',
	"markdown": '0',
	"md":       '0',
	"mp3":      's',
	"png":      'I',
	"sh":       '0',
	"txt":      '0',
	"wav":      's',
}

type Listing struct {
	entries []Entry
}

func (l Listing) String() string {
	var b bytes.Buffer

	for _, e := range l.entries {
		if e.Type == '1' {
			fmt.Fprint(&b, e)
		}
	}

	for _, e := range l.entries {
		if e.Type == 0 || e.Type == '1' {
			continue // skip sentinel value and directories
		}

		fmt.Fprint(&b, e)
	}

	return b.String()
}

func (l *Listing) VisitDir(name, path, root, host, port string) error {
	if len(l.entries) == 0 {
		l.entries = append(l.entries, Entry{}) // sentinel value
		return nil
	}

	l.entries = append(l.entries, Entry{'1', name, path[len(root):], host, port})

	return filepath.SkipDir
}

func (l *Listing) VisitFile(name, path, root, host, port string) {
	t := byte('9') // Binary

	for s, c := range Suffixes {
		if strings.HasSuffix(path, "."+s) {
			t = c
			break
		}
	}

	l.entries = append(l.entries, Entry{t, name, path[len(root):], host, port})
}
