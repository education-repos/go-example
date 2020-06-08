package handlers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// хендлер для обращения на внешний ресурс
func ApiHandler(w http.ResponseWriter, r *http.Request) {
	// формирование запроса
	req := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"Authorization": {prepareAuthorization()},
		},
	}
	req.URL, _ = url.Parse("https://kazanexpress.ru/api/main/root-categories")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error making request")
	}
	defer resp.Body.Close()

	// вывод тела ответа
	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, "%#v\n\n\n", string(respBody))
}

// подготовка заголовка авторизации
func prepareAuthorization() string {
	authType := "Basic "
	authString := []byte("kazanexpress-customer:customerSecretKey")
	return authType + base64.StdEncoding.EncodeToString(authString)
}
