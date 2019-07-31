package main

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var renderer *render.Render

func init() {
	renderer = render.New()
}

func main() {
	router := httprouter.New()
	router.GET("/", IndexHandler)
	router.GET("/index", IndexHandler)
	router.GET("/login", LoginHandler)
	router.POST("/login", LoginPostHandler)
	router.GET("/logout", LogoutHandler)
	n := negroni.New()
	n.UseHandler(router)
	n.Run(":5000")
}

func IndexHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	cookie, err := req.Cookie("jwtToken")
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	}
	err = AuthenticateToken(cookie.Value)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusFound)
		return
	}
	renderer.HTML(w, http.StatusOK, "index", map[string]string{"message": "Hello, World"})
	return
}

func LoginHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	cookie, _ := req.Cookie("jwtToken")
	if cookie != nil {
		err := AuthenticateToken(cookie.Value)
		if err == nil {
			http.Redirect(w, req, "index", http.StatusFound)
			return
		}
	}
	renderer.HTML(w, http.StatusOK, "login", map[string]string{"message": "Login is Required"})
}

func LoginPostHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	expirationTime := time.Now().Add(5 * time.Minute)
	user := User{
		req.FormValue("Username"),
		req.FormValue("Password"),
		req.FormValue("Birthday"),
		req.FormValue("Email"),
	}
	tokenString, err := GenerateToken(user)
	if err != nil {
		renderer.HTML(w, http.StatusOK, "login", map[string]string{"message": "Check your 'Username' or 'Password'"})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "jwtToken",
		Value:   tokenString,
		Expires: expirationTime,
	})
	http.Redirect(w, req, "/index", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	cookie, err := req.Cookie("jwtToken")
	if err != nil {
		http.Redirect(w, req, "login", http.StatusFound)
		return
	}
	//cookie 삭제
	cookie.Expires = time.Now().Add(-100 * time.Minute)
	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/index", http.StatusFound)
}
