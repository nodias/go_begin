package main

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	// 키: http 메서드
	// 값: URL 패턴별로 실행할 HandlerFunc
	handlers map[string]map[string]HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h HandlerFunc) {
	log.Println("# HandleFunc")
	// http 메서드로 등록된 맵이 있는지 확인
	m, ok := r.handlers[method]
	if !ok {
		// 등록된 맵이 없으면 새 맵을 생성
		m = make(map[string]HandlerFunc)
		r.handlers[method] = m
	}
	// http 메서드로 등록된 맵에 URL 패턴과 핸들러 함수 등록
	m[pattern] = h
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// http 메서드에 맞는 모든 handers를 반복하면서 요청 URL에 해당하는 handler를 찾음
	for pattern, handlerF := range r.handlers[req.Method] {
		if ok, params := match(pattern, req.URL.Path); ok {
			context := Context{params, w, req}
			handlerF(&context)
			return
		}
	}
	// 요청 URL에 해당하는 handler 수행
	http.NotFound(w, req)
	return
}

func match(pattern, path string) (bool, map[string]string) {
	if pattern == path {
		return true, nil
	}
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		return false, nil
	}

	params := make(map[string]string)

	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			params[patterns[i][1:]] = paths[i]
		default:
			return false, nil
		}
	}
	return true, params
}
