#!/bin/zsh

# This script exists primary to document the different actions that have been
# taken to build and use this application

source scripts/console.lib.sh

# Compile a version of the app for a specified operating system and chip set
function build_for() {
	cprintln $FYELLOW "Build for $1 on processor $2."
	GOOS=$1 GOARCH=$2 go build -o gopher.$1.$2 cmd/gopherd/main.go
}

# ##############################################################################

# Build all supported binaries
function build_all() {
	cprint $FGREEN "Check that the app passes tests"
	go test
	if [ "$?" = "0" ]
	then
		cprintln $FGREEN "Build all binaries"
		go build -o gopher.local.native cmd/gopherd/main.go
		build_for linux amd64 #intel based linux
		build_for darwin arm64 #m1 based macs
		build_for darwin amd64 #intel based macs
		build_for windows amd64 #legacy junk computers
	else
		cprint $FRED "Test failed:" ; echo "Not building executables"
	fi
}

# Print a helpfull manual on flags
function usage() {
	printf "Usage:\n"
	format="%4s %-4s : %s\n"
	printf "${format}" "Flag" "Arg" "Description"
	printf "${format}" "----" "----" "----------"
	printf "${format}" "h" "" "This message"
	printf "${format}" "b" "" "Build all binaries"
	printf "${format}" "t" "" "Test everything"
	printf "${format}" "v" "" "Vet (lint) everything"
	printf "${format}" "x" "text" "Print a separator"
}

# ##############################################################################

# Behavior for no parameters
if [ $# -eq 0 ]; then
	usage
	exit 1
fi

while getopts bcfhrtvx: opt ; do
    case "${opt}" in
        b) build_all ;;
        c) go build -o gopher.local.native cmd/gopherd/main.go ;;
        f) gofmt -d . ;;
        h) usage ;;
        r) ./gopher.local.native -host localhost -port 7070 -root content ;;
        t) go test ;;
        v) go vet ./... ;;
        x) echo "${OPTARG}" ;;
        *) cprintln $FRED "Unknown flag" ; usage ;;
    esac
done
