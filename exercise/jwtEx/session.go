package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	sessions "github.com/goincremental/negroni-sessions"
)

const (
	currentUserKey  = "oauth2_current_user"
	sessionDuration = time.Hour
)

type User struct {
	Uid       string    `json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"user"`
	AvatarUrl string    `json:"avatar_url"`
	Expired   time.Time `json:"expired"`
}

func (u *User) Vailid() bool {
	return u.Expired.Sub(time.Now()) > 0
}

func (u *User) Refresh() {
	u.Expired = time.Now().Add(sessionDuration)
}

func GetCurrentUser(r *http.Request) (u *User) {
	var session sessions.Session
	session = sessions.GetSession(r)
	data := session.Get(currentUserKey)
	if data != nil {
		json.Unmarshal(data.([]byte), &u)
		return
	}
	return
}

func SetCurrentUser(r *http.Request, u *User) {
	if u != nil {
		u.Refresh()
	}
	session := sessions.GetSession(r)
	val, _ := json.Marshal(u)
	session.Set(currentUserKey, val)
}

func DeleteCurrentUser(r *http.Request, u *User) {
	if u != nil {
		sessions.GetSession(r).Delete(currentUserKey)
		return
	}
	log.Println("There is no user")
	return
}
