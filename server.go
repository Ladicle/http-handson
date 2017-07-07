package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.Int("p", 7070, "server port number")
	host = flag.String("h", "127.0.0.1", "server host name")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)

	log.Printf("Access to %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, GenerateRouter()))
}
