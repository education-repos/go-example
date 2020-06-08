package handlers

import (
	"fmt"
	"net/http"
)

// общий хендлер для всех запросов
func MainHandler(w http.ResponseWriter, r *http.Request) {
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
