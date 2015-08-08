package gopher

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

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

func (s *Server) Serve(c *net.TCPConn) {
	defer c.Close()

	p, err := s.getPath(c)
	if err != nil {
		fmt.Fprint(c, s.Error("invalid request"))

		return
	}

	fn := s.filename(p)

	fi, err := os.Stat(fn)
	if err != nil {
		fmt.Fprint(c, s.Error("not found"))

		return
	}

	if fi.IsDir() {
		var list Listing

		filepath.Walk(fn, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return list.VisitDir(path, s.Root, s.Host, s.Port, info)
			}

			list.VisitFile(path, s.Root, s.Host, s.Port, info)

			return nil
		})

		fmt.Fprint(c, list)

		return
	}

	f, err := os.Open(fn)
	if err != nil {
		fmt.Fprint(c, s.Error("couldn't open file"))
		return
	}
	defer f.Close()

	c.ReadFrom(f)
}

func (s *Server) Error(msg string) Listing {
	return Listing{[]Entry{{Type: 3, Display: msg}}}
}

func (s *Server) filename(p []byte) string {
	return s.Root + filepath.Clean("/"+string(p))
}

func (s *Server) getPath(c *net.TCPConn) ([]byte, error) {
	p, _, err := bufio.NewReader(c).ReadLine()

	return p, err
}
