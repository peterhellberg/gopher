package gopher

import (
	"bytes"
	"fmt"
	"os"
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
		if e.Type == 0 {
			continue // skip sentinel value
		}

		if e.Type == 1 {
			fmt.Fprint(&b, e)
		}
	}

	for _, e := range l.entries {
		if e.Type == 0 {
			continue // skip sentinel value
		}

		fmt.Fprint(&b, e)
	}

	return b.String()
}

func (l *Listing) VisitDir(path, root, host, port string, f os.FileInfo) error {
	if len(l.entries) == 0 {
		l.entries = append(l.entries, Entry{}) // sentinel value
		return nil
	}

	l.entries = append(l.entries, Entry{'1', f.Name(), path[len(root):], host, port})

	return filepath.SkipDir
}

func (l *Listing) VisitFile(path, root, host, port string, f os.FileInfo) {
	t := byte('9') // Binary

	for s, c := range Suffixes {
		if strings.HasSuffix(path, "."+s) {
			t = c
			break
		}
	}

	l.entries = append(l.entries, Entry{t, f.Name(), path[len(root):], host, port})
}
