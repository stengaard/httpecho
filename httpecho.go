// Command httpecho simply returns a copy of the incoming request over the channel it is listening on.
package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func Echo(w http.ResponseWriter, req *http.Request) {
	spew.Fdump(w, req)
}

func main() {
	addr := flag.String("bind", "", "The address to bind to")

	flag.Parse()
	var l net.Listener
	var err error

	// gin support
	if *addr == "" {
		*addr = fmt.Sprintf("tcp://:%s/", os.Getenv("PORT"))
	}

	u, err := url.Parse(*addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parse bind url: %s\n", err)
	}

	if u.Host == "" {
		u.Host = u.Path
	}

	l, err = net.Listen(u.Scheme, u.Host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err listening: %s\n", err)
		os.Exit(1)
	}

	if err := http.Serve(l, http.HandlerFunc(Echo)); err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s\n", err)
		os.Exit(1)
	}
}
