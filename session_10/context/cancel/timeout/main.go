package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hi/timeout", handlerTimeout)
	http.ListenAndServe(":9090", nil)
}

func handlerTimeout(w http.ResponseWriter, r *http.Request){
	toCtx, cancel := context.WithTimeout(r.Context(), 3 * time.Second)
	defer cancel()
	working := make(chan string)
	go func(ctx context.Context, w chan<- string) {
		select {
		case <-time.After(2 * time.Second):
			w <- "finished1"
			case <-ctx.Done():
				fmt.Printf("\n canceled into goroutine1 %s", time.Now().String())
		}
	}(toCtx, working)
	go func(ctx context.Context, w chan<- string) {
		select {
		case <-time.After(4 * time.Second):
			w <- "finished2"
		case <-ctx.Done():
			fmt.Printf("\n canceled into goroutine2 %s", time.Now().String())
		}
	}(toCtx, working)
	select {
	case message := <-working:
		fmt.Fprintln(w, message)
	case <-toCtx.Done():
		fmt.Printf("\n cancelling %s %s", toCtx.Err().Error(), time.Now().String())
		cancel()
		http.Error(w, toCtx.Err().Error(), http.StatusInternalServerError)
	}
}

