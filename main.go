package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_server/internal/configs"
	"go_server/internal/router"
	"net/http"
	"time"
)

func main() {
	var cfg configs.Config
	configs.ReadFile(&cfg)
	//readEnv(&cfg)

	mux := http.NewServeMux()
	router.RegisterRoutes(mux)

	server := http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at :8080")
	fmt.Println("http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error starting server at :8080")
	}
}
