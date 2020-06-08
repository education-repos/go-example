package handlers

import "net/http"

// хендлер для обработки статики
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix(
		"/data/",
		http.FileServer(http.Dir("./static")),
	)
}
