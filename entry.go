package gopher

import "fmt"

// Entry for use in Gopher listings
type Entry struct {
	Type     byte
	Display  string
	Selector string
	Hostname string
	Port     int
}

// String returns entry as a Gopher listing formatted string
func (e Entry) String() string {
	return fmt.Sprintf("%c%s\t%s\t%s\t%d\r\n",
		e.Type, e.Display, e.Selector, e.Hostname, e.Port)
}
