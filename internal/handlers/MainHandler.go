package handlers

import (
	"fmt"
	"io"
	"net/http"

	"shortener/internal/utils"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// Закрываем тело запроса
		defer r.Body.Close()

		randStr := "/" + utils.RandSeq(10)
		new_url := fmt.Sprintf("http://localhost:8080%s", randStr)

		utils.UrlsMap[randStr] = string(body)

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, new_url)
	}
	if r.Method == http.MethodGet {

		url := r.URL.String()

		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", utils.UrlsMap[url])
		w.WriteHeader(http.StatusTemporaryRedirect)
	}

}

func urlPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		url := r.URL.String()

		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Location", utils.UrlsMap[url])
		w.WriteHeader(http.StatusTemporaryRedirect)
	}

}
