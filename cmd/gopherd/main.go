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
	port = flag.Int("port", 7070, "listen on port")
	root = flag.String("root", "", "root directory for content")
	once = flag.String("once", "", "process one file and exit, CGI mode")
)

func main() {
	parseFlags()

	// Setup the logger used by the server
	logger := log.New(os.Stdout, "", 0)

	// Combine host and port into addr
	addr := net.JoinHostPort(*host, fmt.Sprint(*port))

	// Create the server
	server := &gopher.Server{
		Logger: logger,
		Host:   *host,
		Port:   *port,
		Addr:   addr,
		Root:   *root,
	}

	if *once != "" {
		//get content and return it to the display
		result := server.ServeFile(os.Stdout, *once)
		fmt.Print(result)
		return
	}

	server.Logger.Printf("Starting server...")
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

	normalized_root := strings.TrimSuffix(*root, "/")
	root = &normalized_root

	if *root == "" || *root == "." {
		dir, err := os.Getwd()
		if err != nil {
			flag.Usage()
		}
		root = &dir
	}
}
