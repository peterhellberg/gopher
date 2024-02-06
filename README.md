# gopher

[![Build Status](https://travis-ci.org/peterhellberg/gopher.svg?branch=master)](https://travis-ci.org/peterhellberg/gopher)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/gopher)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/gopher#license-mit)

A server for the [Gopher](https://en.wikipedia.org/wiki/Gopher_%28protocol%29)
protocol written in [Go](http://golang.org/) initially based on [gogopherd](https://github.com/nf/gogopherd) by [Andrew Gerrand](https://twitter.com/enneff).

> **Note:** This project is meant to be an experiment, do not expect any stability guarantees.

## Installation

    go get -u github.com/peterhellberg/gopher/cmd/gopherd

## Compile

	go build -o gopher.local.native cmd/gopherd/main.go

## Usage
For a standalone server use the following command to run the server on the localhost under port 7070.

	./gopher.local.native \
		-host localhost \
		-port 7070 \
		-root content

For usage as a CGI script or to just test one call, try the `-once` flag with the request of interest. Provide a `-host` and `-port` as normal because these values are used in the return content.

	./gopher.local.native \
	    -host localhost \
	    -port 7070 \
	    -root content \
	    -once /README.md

## License (MIT)

Copyright (c) 2015-2018 [Peter Hellberg](https://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

<img src="https://data.gopher.se/gopher/viking-gopher.svg" align="right" width="230" height="230">

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
