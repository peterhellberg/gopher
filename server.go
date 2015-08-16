package gopher

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

// Server is a Gopher server
type Server struct {
	Logger *log.Logger
	Host   string
	Port   string
	Addr   string
	Root   string
}

// ListenAndServe starts listening
func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", s.Host+":"+s.Port)
	if err != nil {
		return err
	}

	s.Logger.Printf("Listening on gopher://%s:%s\n",
		s.Host, s.Port)

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}

		go s.Serve(c.(*net.TCPConn))
	}
}

// Serve Gopher clients
func (s *Server) Serve(c *net.TCPConn) {
	defer c.Close()

	p, err := s.getPath(c)
	if err != nil {
		fmt.Fprint(c, s.ErrorListing("invalid request"))

		return
	}

	fn := s.filename(p)

	fi, err := os.Stat(fn)
	if err != nil {
		fmt.Fprint(c, s.ErrorListing("not found"))

		return
	}

	if fi.IsDir() {
		var list Listing

		filepath.Walk(fn, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return list.VisitDir(info.Name(), path, s.Root, s.Host, s.Port)
			}

			list.VisitFile(info.Name(), path, s.Root, s.Host, s.Port)

			return nil
		})

		fmt.Fprint(c, list)

		return
	}

	f, err := os.Open(fn)
	if err != nil {
		fmt.Fprint(c, s.ErrorListing("couldn't open file"))
		return
	}
	defer f.Close()

	c.ReadFrom(f)
}

// ErrorListing creates an error listing
func (s *Server) ErrorListing(msg string) Listing {
	return Listing{[]Entry{{Type: 3, Display: msg}}}
}

func (s *Server) filename(p string) string {
	return s.Root + filepath.Clean("/"+p)
}

func (s *Server) getPath(rd io.Reader) (string, error) {
	p, _, err := bufio.NewReader(rd).ReadLine()

	return string(p), err
}
