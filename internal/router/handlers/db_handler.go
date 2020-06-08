package handlers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go_server/internal/configs"
	"go_server/internal/entities"
	"net/http"
)

// хендлер для обращения в БД
func DbHandler(w http.ResponseWriter, r *http.Request) {
	var cfg configs.Config
	configs.ReadFile(&cfg)
	conn, err := sqlx.Open(cfg.Database.Driver, configs.DsnString(&cfg))
	if err != nil {
		fmt.Println("unable to open connect to db")
	}
	var category entities.Category
	id := 3
	err = conn.Get(&category, "select * from category where id=?", id)
	if err != nil {
		panic(err)
	}
	// вывод тела ответа
	fmt.Fprintf(w, "%#v\n\n\n", category)
}
