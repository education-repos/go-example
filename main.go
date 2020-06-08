package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// общий хендлер для всех запросов
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

// хендлер для вывода изображения
func imgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<img src='/data/img/go_logo.png' />`))
}

// хендлер для обращения на внешний ресурс
func apiHandler(w http.ResponseWriter, r *http.Request) {
	// подготовка заголовка авторизации
	authType := "Basic "
	authString := []byte("kazanexpress-customer:customerSecretKey")
	decodedAuthString := authType + base64.StdEncoding.EncodeToString(authString)

	// формирование запроса
	req := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"Authorization": {decodedAuthString},
		},
	}
	req.URL, _ = url.Parse("https://kazanexpress.ru/api/main/root-categories")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.Write([]byte("Error"))
	}
	defer resp.Body.Close()

	// вывод тела ответа
	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, "%#v\n\n\n", string(respBody))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/image", imgHandler)
	mux.HandleFunc("/api", apiHandler)

	// обработка статики
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
