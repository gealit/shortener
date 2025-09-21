package handlers

import "net/http"

func Route() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(`/`, mainPage)
	// mux.HandleFunc(`/`, urlPage)

	return mux
}
