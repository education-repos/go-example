package router

import (
	"go_server/internal/router/handlers"
	"net/http"
)

// роутинг
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/image", handlers.ImgHandler)
	mux.HandleFunc("/api", handlers.ApiHandler)
	mux.HandleFunc("/db", handlers.DbHandler)
	mux.HandleFunc("/data/", handlers.StaticHandler)
}
