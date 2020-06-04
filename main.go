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

	ua := r.UserAgent()
	if ua != "" {
		fmt.Fprintln(w, "Your User Agent is "+ua)
	}

	accept := r.Header.Get("Accept")
	if accept != "" {
		fmt.Fprintln(w, "You accept "+accept)
	}

	w.Write([]byte("======================"))
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<img src='/data/img/go_logo.png' />`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/image", imgHandler)

	staticHandler := http.StripPrefix(
		"/data/",
		http.FileServer(http.Dir("./static")),
	)

	mux.Handle("/data/", staticHandler)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at :8080")
	server.ListenAndServe()
}
