package main

import (
	"encoding/json"
	"fmt"
	"github.com/nodias/go_begin/book_buildingmicroservicesGO/common"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string
}

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello, World"}
	data := common.MustByte(json.Marshal(response))
	fmt.Fprint(w, string(data))
}


