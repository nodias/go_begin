package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	r := &router{make(map[string]map[string]HandlerFunc)}

	r.HandleFunc("GET", "/", func(c *Context) {
		t := time.Now()
		fmt.Fprintln(c.ResponseWriter, "index")
		log.Printf("[%s] %q %v\n", c.Request.Method, c.Request.URL.String(), time.Now().Sub(t))
	})

	r.HandleFunc("GET", "/about", func(c *Context) {
		t := time.Now()
		fmt.Fprintln(c.ResponseWriter, "about")
		log.Printf("[%s] %q %v\n", c.Request.Method, c.Request.URL.String(), time.Now().Sub(t))
	})

	r.HandleFunc("GET", "/users/:id", func(c *Context) {
		t := time.Now()
		fmt.Fprintln(c.ResponseWriter, "retrieve user", c.Params["id"])
		log.Printf("[%s] %q %v\n", c.Request.Method, c.Request.URL.String(), time.Now().Sub(t))
	})

	r.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		t := time.Now()
		fmt.Fprintln(c.ResponseWriter, "retrieve user’s address")
		log.Printf("[%s] %q %v\n", c.Request.Method, c.Request.URL.String(), time.Now().Sub(t))
	})

	r.HandleFunc("POST", "/users", func(c *Context) {
		t := time.Now()
		fmt.Fprintln(c.ResponseWriter, "create user")
		log.Printf("[%s] %q %v\n", c.Request.Method, c.Request.URL.String(), time.Now().Sub(t))
	})

	r.HandleFunc("POST", "/users/:user_id/addresses", func(c *Context) {
		t := time.Now()
		fmt.Fprintln(c.ResponseWriter, "create user’s address")
		log.Printf("[%s] %q %v\n", c.Request.Method, c.Request.URL.String(), time.Now().Sub(t))
	})

	http.ListenAndServe(":8080", r) //(addr string, Handler interface)

	/**
	type Handler interface{
		ServeHTTP(http.ResponseWriter, *http.Request)
	}
	*/
}
