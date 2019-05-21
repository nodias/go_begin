package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Start")
	r := &router{make(map[string]map[string]HandlerFunc)}

	r.HandleFunc("GET", "/", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "index")
	})))))

	r.HandleFunc("GET", "/about", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "about")
	})))))

	r.HandleFunc("GET", "/users/:id", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "retrieve user", c.Params["id"])
	})))))

	r.HandleFunc("GET", "/users/:user_id/addresses/:address_id", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "retrieve user’s address")
	})))))

	r.HandleFunc("POST", "/users", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "create user")
	})))))

	r.HandleFunc("POST", "/users/:user_id/addresses", logHandler(recoverHandler(parseFormHandler(parseJsonBodyHandler(func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "create user’s address")
	})))))

	http.ListenAndServe(":8080", r) //(addr string, Handler interface)

	/**
	type Handler interface{
		ServeHTTP(http.ResponseWriter, *http.Request)
	}
	*/
}
