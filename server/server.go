package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		//w.Header().Set("Connection", "Close")
		fmt.Fprintf(w, "hello\n")
	})

	srv := &http.Server{
		Addr:        ":8090",
		Handler:     mux,
		IdleTimeout: time.Millisecond * 200, // to cause EOF
	}
	srv.ListenAndServe()
}
