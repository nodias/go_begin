package main

import (
	"log"
	"net/http"

	"github.com/goincremental/negroni-sessions/cookiestore"

	sessions "github.com/goincremental/negroni-sessions"

	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var renderer *render.Render

const (
	sessionKey    = "mysession"
	sessionSecret = "mysession_secret"
)

func init() {
	log.Println("### init server")
	renderer = render.New()
}

func main() {
	log.Println("### main")

	//router
	router := httprouter.New()
	router.GET("/", IndexHandler)
	router.GET("/index", IndexHandler)
	router.GET("/login", LoginHandler)
	router.GET("/logout", LogoutHandler)

	//negroni
	n := negroni.Classic()

	//session
	store := cookiestore.New([]byte(sessionSecret))
	n.Use(sessions.Sessions(sessionKey, store))

	n.UseHandler(router)
	n.Run(":5000")
}

func IndexHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("### IndexHandler")
	session := sessions.GetSession(req)
	sessionVal := session.Get("sessionkey")
	if sessionVal != nil {
		renderer.HTML(w, http.StatusOK, "index", map[string]string{"message": sessionVal.(string)})
		return
	}
	renderer.HTML(w, http.StatusOK, "index", map[string]string{"message": "please login"})
}

func LoginHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("### LoginHandler")
	session := sessions.GetSession(req)
	session.Set("sessionkey", "sessionval")
	http.Redirect(w, req, "/index", http.StatusFound)
}

func LogoutHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	log.Println("### LogoutHandler")
	session := sessions.GetSession(req)
	session.Delete("sessionkey")
	http.Redirect(w, req, "/index", http.StatusFound)
}
