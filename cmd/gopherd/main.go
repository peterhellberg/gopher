package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/peterhellberg/gopher"
)

var (
	host = flag.String("host", "0.0.0.0", "hostname used in links")
	port = flag.String("port", "7070", "listen on port")

	root string
)

func main() {
	parseFlags()

	// Setup the logger used by the server
	logger := log.New(os.Stdout, "", 0)

	// Combine host and port into addr
	addr := net.JoinHostPort(*host, *port)

	// Create the server
	server := &gopher.Server{
		Logger: logger,
		Host:   *host,
		Port:   *port,
		Addr:   addr,
		Root:   root,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func parseFlags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s directory\n", os.Args[0])

		flag.PrintDefaults()

		os.Exit(2)
	}

	flag.Parse()

	if root = strings.TrimSuffix(flag.Arg(0), "/"); root == "" {
		flag.Usage()
	}
}
