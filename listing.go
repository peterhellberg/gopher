package gopher

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
)

// Suffixes is a map of file extensions to item type characters
var Suffixes = map[string]byte{
	"aiff":     's',
	"au":       's',
	"c":        '0',
	"cfg":      '0',
	"cpp":      '0',
	"cs":       '0',
	"css":      '0',
	"csv":      '0',
	"gif":      'g',
	"go":       '0',
	"gpg":      '0',
	"h":        '0',
	"html":     'h',
	"ini":      '0',
	"java":     '0',
	"jpeg":     'I',
	"jpg":      'I',
	"js":       '0',
	"json":     '0',
	"log":      '0',
	"lua":      '0',
	"markdown": '0',
	"md":       '0',
	"mp3":      's',
	"php":      '0',
	"pl":       '0',
	"png":      'I',
	"py":       '0',
	"rb":       '0',
	"rss":      '0',
	"sh":       '0',
	"txt":      '0',
	"wav":      's',
	"xml":      '0',
}

// Listing is a Gopher listing containing entries
type Listing struct {
	entries []Entry
}

// String returns a Gopher listing formatted string
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

// VisitDir appends a dir entry to the list of entries in the listing
func (l *Listing) VisitDir(name, path, root, host string, port int) error {
	if len(l.entries) == 0 {
		l.entries = append(l.entries, Entry{}) // sentinel value
		return nil
	}

	l.entries = append(l.entries, Entry{'1', name, path[len(root):], host, port})

	return filepath.SkipDir
}

// VisitFile appends a file entry to the list of entries in the listing
func (l *Listing) VisitFile(name, path, root, host string, port int) {
	t := byte('9') // Binary

	for s, c := range Suffixes {
		if strings.HasSuffix(path, "."+s) {
			t = c
			break
		}
	}

	l.entries = append(l.entries, Entry{t, name, path[len(root):], host, port})
}
