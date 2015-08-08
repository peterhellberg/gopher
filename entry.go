package gopher

import "fmt"

type Entry struct {
	Type     byte
	Display  string
	Selector string
	Hostname string
	Port     string
}

func (e Entry) String() string {
	return fmt.Sprintf("%c%s\t%s\t%s\t%s\r\n",
		e.Type, e.Display, e.Selector, e.Hostname, e.Port)
}
