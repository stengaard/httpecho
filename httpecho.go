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
	die(err)

	if u.Scheme == "unix" {
		l, err = net.Listen(u.Scheme, u.Path)
		die(err)
		die(os.Chmod(u.Path, 0666))
	} else {
		l, err = net.Listen(u.Scheme, u.Host)
		die(err)
	}

	die(http.Serve(l, http.HandlerFunc(Echo)))
}

func die(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s\n", err)
		os.Exit(1)
	}
}
