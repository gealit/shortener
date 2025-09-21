package main

import (
	"net/http"

	"shortener/internal/handlers"

	"shortener/internal/utils"
)

func main() {

	utils.UrlsMap = make(map[string]string)

	router := handlers.Route()

	err := http.ListenAndServe(`:8080`, router)
	if err != nil {
		panic(err)
	}
}
