package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang/glog"
)

var (
	port = flag.Int("port", 8080, "Port on which to serve")
)

func main() {
	flag.Parse()
	message := os.Getenv("MESSAGE")
	if message == "" {
		glog.Fatal("MESSAGE env var is not set")
	}
	glog.Infof("Starting the test app with message: %q", message)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		glog.Infof("Serving a request")
		fmt.Fprintf(w, "The server greets you: %s\r\n", message)
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
