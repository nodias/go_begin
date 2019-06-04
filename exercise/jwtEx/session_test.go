package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/urfave/negroni"

	Cont "github.com/nodias/go_begin/exercise/sessionEx"
)

var user = Cont.User{
	Uid:       "uid",
	Name:      "name",
	Email:     "Email",
	AvatarUrl: "AvatarUrl",
	Expired:   time.Now().Add(time.Hour),
}

func TestGetCurrentUser(t *testing.T) {
	n := negroni.Classic()

	store := cookiestore.New([]byte("secret123"))
	n.Use(sessions.Sessions("oauth2_current_user", store))

	mux := http.NewServeMux()
	mux.HandleFunc("/testsession/SetCurrentUser", func(w http.ResponseWriter, req *http.Request) {
		Cont.SetCurrentUser(req, &user)
		fmt.Fprintf(w, "OK")
	})

	mux.HandleFunc("/testsession/GetCurrentUser", func(w http.ResponseWriter, req *http.Request) {
		u := Cont.GetCurrentUser(req)
		fmt.Println("Success : ", u)
	})

	n.UseHandler(mux)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testsession/SetCurrentUser", nil)
	n.ServeHTTP(res, req)

	res2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/testsession/GetCurrentUser", nil)
	req2.Header.Set("Cookie", res.Header().Get("Set-Cookie"))
	n.ServeHTTP(res2, req2)

}
