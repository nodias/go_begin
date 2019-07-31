// Binary fileserver is a file server example.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "address of the webserver")
	root = flag.String("root", "/var/www", "root directory")
)

func main() {
	flag.Parse()
	fmt.Printf("%s\n%s", *addr, *root)
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
