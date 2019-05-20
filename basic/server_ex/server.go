package main

import (
	"log"
	"net/http"

	io "github.com/nodias/go_begin/basic/server_ex/io"
)

func main() {
	log.Println("Server Starts : 5000")
	indexPage := func(res http.ResponseWriter, req *http.Request) {
		indexHtmlfile := "index.html"
		indexHTML := io.ReadHtml(indexHtmlfile)

		res.Header().Set("Content-Type", "text/html")
		res.Write(indexHTML)
	}
	http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", indexPage)
	http.ListenAndServe(":5000", nil)
}
