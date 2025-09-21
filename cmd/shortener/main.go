package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gealit/shortener/internal/service"
)

var UrlsMap map[string]string

func Route() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(`/`, MainPage)
	mux.HandleFunc(`/EwHXdJfB`, urlPage)

	return mux
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// Закрываем тело запроса
		defer r.Body.Close()

		randStr := "/" + service.RandSeq(10)
		new_url := fmt.Sprintf("http://localhost:8080%s", randStr)

		UrlsMap[randStr] = string(body)

		fmt.Println("В ulrsMap записно:", randStr, UrlsMap[randStr])

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		// fmt.Fprint(w, string(body))
		fmt.Fprint(w, new_url)
	}
	if r.Method == http.MethodGet {

		url := r.URL.String()
		fmt.Println("Получен URL:", url)

		// HTTP/1.1 307 Temporary Redirect
		// Location: https://practicum.yandex.ru/
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", UrlsMap[url])
		w.WriteHeader(http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "Location: %s", UrlsMap[url])
	}

}

func urlPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		url := r.URL.String()
		fmt.Println("Получен URL:", url)

		// HTTP/1.1 307 Temporary Redirect
		// Location: https://practicum.yandex.ru/
		w.WriteHeader(http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "Location: %s", UrlsMap[url])
	}

}

func main() {

	UrlsMap = make(map[string]string)

	router := Route()

	err := http.ListenAndServe(`:8080`, router)
	if err != nil {
		panic(err)
	}
}
