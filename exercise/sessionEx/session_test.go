package main

import (
	"net/http"
	"testing"
)

func TestGetCurrentUser(t *testing.T) {
	var req = &http.Request{}

	GetCurrentUser(req)

}
