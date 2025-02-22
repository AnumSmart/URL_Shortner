package link

import (
	"fmt"
	"net/http"
	"server/pkg/middleware"
	"server/pkg/req"
	"server/pkg/resp"
	"strconv"

	"gorm.io/gorm"
)

type LinkHandlerDep struct {
	LinkRepository *LinkRepository
}

type LinkHandler struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDep) {
	handler := &LinkHandler{
		LinkRepository: deps.LinkRepository,
	}
	router.HandleFunc("/link/create", handler.Create())
	router.HandleFunc("/link/{hash}", handler.GoTo())
	router.Handle("/link/{id}", middleware.IsAuthed(handler.Update()))
	router.HandleFunc("/link/{id}", handler.Delete())
}

func (l *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "create link")

		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		//-------------------------------------------------------------
		//check that we dont have the same HASH in DB
		link := NewLink(body.Url)
		for {
			existedLink, _ := l.LinkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}
			link.GenerateHash()
		}
		//-------------------------------------------------------------
		createdLink, err := l.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp.ResponceJson(w, 201, createdLink)
	}
}

func (l *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GoTo link")
		hash := r.PathValue("hash")
		link, err := l.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (l *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "update link")
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}

		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := l.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp.ResponceJson(w, 201, link)
	}
}

func (l *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "delete link")
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//-----------------------------------------------------
		// проверка наличия элемента, который хотим удалить
		_, err = l.LinkRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		//-----------------------------------------------------
		err = l.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp.ResponceJson(w, 200, nil)
	}
}
