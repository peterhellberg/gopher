package gopher

import (
	"bufio"
	"bytes"
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
	Port   int
	Addr   string
	Root   string
}

// ListenAndServe starts listening
func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		return err
	}

	s.Logger.Printf("Listening on gopher://%s:%d and serving files from %s.\n",
		s.Host, s.Port, s.Root)

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
		s.Logger.Print(s.ErrorListing("invalid request"))

		return
	}

	s.Logger.Printf("Request for %s.", p)
	output := s.ServeFile(c, p)
	fmt.Fprint(c, output)
}

func (s *Server) ServeFile(c io.Writer, p string) string {
	fn := s.filename(p)

	fi, err := os.Stat(fn)
	if err != nil {
		fmt.Fprint(c, s.ErrorListing("not found"))
		//no output, it's an error
		return "3'Error'\tFile not found\terror.host\t1"
	}

	//directories
	if fi.IsDir() {
		var list Listing

		filepath.Walk(fn, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return list.VisitDir(info.Name(), path, s.Root, s.Host, s.Port)
			}
			list.VisitFile(info.Name(), path, s.Root, s.Host, s.Port)

			return nil
		})
		return list.String()
	}

	//files
	f, err := os.Open(fn)
	if err != nil {
		s.Logger.Print(s.ErrorListing("couldn't open file"))
		return "3'Error'\tCould not open file\terror.host\t1"
	}
	defer f.Close()

	buffer := bytes.Buffer{}
	_, err = buffer.ReadFrom(f)
	if err != nil {
		s.Logger.Print(err)
		return "3'Error'\tError Reading File\terror.host\t1"
	}
	return buffer.String()
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
