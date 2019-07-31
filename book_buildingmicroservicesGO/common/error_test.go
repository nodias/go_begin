package common

import (
	"errors"
	"reflect"
	"testing"
)

func TestMustByte(t *testing.T) {
	type params struct {
		b       []byte
		err     error
		ifpanic func(d string, p params)
	}
	var byteMock = []byte("bytemock")
	var errorMock = errors.New("errorMock")
	var panicMock = func(d string, p params) {
		if r := recover(); r != nil {
			if !reflect.DeepEqual(p.err, r) {
				t.Errorf("%s: %s != %s", d, p.err, r)
			}
		}
	}
	examples := []struct {
		desc     string
		input    params
		expected []byte
	}{
		{
			desc: "success",
			input: params{
				byteMock,
				nil,
				nil,
			},
			expected: byteMock,
		},
		{
			desc: "panic case",
			input: params{
				nil,
				errorMock,
				panicMock,
			},
			expected: nil,
		},
	}
	for _, ex := range examples {
		if ex.input.ifpanic != nil {
			defer ex.input.ifpanic(ex.desc, ex.input)
		}
		actual := MustByte(ex.input.b, ex.input.err)
		if !reflect.DeepEqual(ex.expected, actual) {
			t.Errorf("%s: %s != %s", ex.desc, ex.expected, actual)
		}
	}
}
