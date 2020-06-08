package handlers

import "net/http"

// хендлер для вывода изображения
func ImgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<img src='/data/img/go_logo.png' />`))
}
