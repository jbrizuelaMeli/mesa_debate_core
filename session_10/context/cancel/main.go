package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hi", handlerHola)
	http.ListenAndServe(":9090", nil)
}

func handlerHola(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	select {
		case <- time.After(2 * time.Second):
			fmt.Println("working")
			fmt.Fprintln(writer, "finished")
		case <-ctx.Done():
			fmt.Printf("canceled: %s \n", request.Context().Err().Error())
	}
}