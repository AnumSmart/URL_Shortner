package link

import (
	"fmt"
	"net/http"
)

type LinkHandlerDep struct {
}

type LinkHandler struct {
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDep) {
	handler := &LinkHandler{}
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (l *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "create link")
	}
}

func (l *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "update link")
	}
}

func (l *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "read link")
	}
}

func (l *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}
