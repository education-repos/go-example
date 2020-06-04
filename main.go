package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	if param != "" {
		fmt.Fprintln(w, "Param from URL is "+param)
	}

	key := r.FormValue("key")
	if key != "" {
		fmt.Fprintln(w, "Param from URL by FormValue is "+key)
	}

	fmt.Fprintln(w, "test")
	w.Write([]byte("!!!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}
