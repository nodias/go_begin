package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type status int

func (s status) MarshalJSON() ([]byte, error) {
	switch s {
	case VIP:
		return []byte(`"VIP"`), nil
	case FRIENDS:
		return []byte(`"FRIENDS"`), nil
	case BEGINNER:
		return []byte(`"BEGINNER"`), nil
	default:
		return []byte(`"null"`), NotClient
	}
}

func (s *status) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `"VIP"`:
		*s = VIP
		return nil
	case `"FRIENDS"`:
		*s = FRIENDS
		return nil
	case `"BEGINNER"`:
		*s = BEGINNER
		return nil
	default:
		return NotClient
	}
}

type ResponseErr struct {
	Err error `json:"err"`
}

func (r ResponseErr) MarshalJSON() ([]byte, error) {
	if r.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, r.Err)), nil
}

func (r *ResponseErr) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	if v == nil {
		r.Err = nil
		return nil
	}
	switch tv := v.(type) {
	case string:
		if tv == Cheated.Error() {
			r.Err = Cheated
			return nil
		}
		r.Err = errors.New(tv)
		return nil
	default:
		return errors.New("ResponseErr unmarshal failed")
	}
}

var Cheated = errors.New("cheating detected")
var NotClient = errors.New("not my client")

const (
	VIP status = iota
	FRIENDS
	BEGINNER
)

type Person struct {
	Id     int64       `json:"id,string"`
	Email  string      `json:"email"`
	Status status      `json:"status"`
	Err    ResponseErr `json:"err"`
}

func (p *Person) Cheat() {
	p.Err = ResponseErr{Cheated}
}

//test
type Person2 struct {
	Err error `json:"err"`
}

func (p *Person2) Cheat() {
	p.Err = Cheated
}
