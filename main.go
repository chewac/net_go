package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/apex/gateway"
)

func main() {
	port := flag.Int("port", -1, "443")
	flag.Parse()
	listener := gateway.ListenAndServe
	portStr := "n/a"
	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.Handle("/", http.FileServer(http.Dir("./public")))
	}

	log.Fatal(listener(portStr, nil))
}
