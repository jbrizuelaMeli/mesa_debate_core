package main

import (
	"context"
	"fmt"
	"net/http"
)

type key string

const (
	keySay key = "keySay"
	keySay2 key = "keySay2"
)

func main() {
	http.HandleFunc("/hi/say", middleware(handlerSay))
	http.ListenAndServe(":9090", nil)
}

func handlerSay(writer http.ResponseWriter, request *http.Request) {
	value := request.Context().Value(keySay).(string)
	value2 := request.Context().Value(keySay2).(string)
	fmt.Fprintln(writer, value + "  " + value2)
}

func middleware(next http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		valueCtx := context.WithValue(r.Context(), keySay, "Hi Luis")
		valueCtx1 := context.WithValue(valueCtx, keySay2, "/ Bye")
		next.ServeHTTP(w, r.WithContext(valueCtx1))
	}
}