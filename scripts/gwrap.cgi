#!/bin/bash

#This is an example CGI script which can be made executable to Apache for running
#the gopher server one page at a time

if [ "$QUERY_STRING" = "" ] ; then
    QUERY_STRING='/'
fi

if [ "$HTTPS" = "on" ] ; then
    http="https"
else
    http="http"
fi

printf "Content-Type: text/text\n"
printf "X-Script-Name: $SCRIPT_NAME\n"
printf "X-Query-Name: $QUERY_STRING\n"
printf "X-Powered-By: CGI Gopher\n"
printf "\n"

./gopher \
	-host "$http://$HTTP_HOST$SCRIPT_NAME?" \
	-port $SERVER_PORT \
	-root ../gopher \
	-once "$QUERY_STRING"
