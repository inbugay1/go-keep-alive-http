package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello\n")
	})

	srv := &http.Server{
		Addr:        ":8090",
		Handler:     mux,
		IdleTimeout: time.Millisecond * 200,
	}
	srv.ListenAndServe()
}
